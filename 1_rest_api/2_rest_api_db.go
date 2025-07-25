package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Product represents the product entity
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// ProductRepository defines the data access layer interface
type ProductRepository interface {
	Create(p *Product) (int, error)
	GetByID(id int) (*Product, error)
	Update(p *Product) error
	Delete(id int) error
	GetAll() ([]Product, error)
	SearchByName(name string) ([]Product, error)
}

// PostgresProductRepository is a PostgreSQL implementation of ProductRepository
type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(p *Product) (int, error) {
	var id int
	query := `INSERT INTO products (name, price, category) VALUES ($1, $2, $3) RETURNING product_id`
	err := r.db.QueryRow(query, p.Name, p.Price, p.Category).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresProductRepository) GetByID(id int) (*Product, error) {
	p := &Product{}
	query := `SELECT product_id, name, price, category FROM products WHERE product_id = $1`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Category)
	if err == sql.ErrNoRows {
		return nil, nil // no product found
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PostgresProductRepository) Update(p *Product) error {
	query := `UPDATE products SET name = $1, price = $2, category = $3 WHERE product_id = $4`
	result, err := r.db.Exec(query, p.Name, p.Price, p.Category, p.ID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PostgresProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE product_id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *PostgresProductRepository) GetAll() ([]Product, error) {
	query := `SELECT product_id, name, price, category FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *PostgresProductRepository) SearchByName(name string) ([]Product, error) {
	query := `SELECT product_id, name, price, category FROM products WHERE name ILIKE $1`
	likePattern := "%" + name + "%"
	rows, err := r.db.Query(query, likePattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// ProductService defines business logic layer interface
type ProductService interface {
	CreateProduct(p *Product) (*Product, error)
	GetProductByID(id int) (*Product, error)
	UpdateProduct(id int, updated *Product) (*Product, error)
	DeleteProduct(id int) error
	GetAllProducts() ([]Product, error)
	SearchProductsByName(name string) ([]Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(p *Product) (*Product, error) {
	if p.Name == "" {
		return nil, errors.New("product name is required")
	}
	if p.Price < 0 {
		return nil, errors.New("product price cannot be negative")
	}
	id, err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

func (s *productService) GetProductByID(id int) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(id int, updated *Product) (*Product, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, nil
	}

	// Update fields if provided
	if updated.Name != "" {
		existing.Name = updated.Name
	}
	if updated.Price > 0 {
		existing.Price = updated.Price
	}
	if updated.Category != "" {
		existing.Category = updated.Category
	}

	err = s.repo.Update(existing)
	if err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

func (s *productService) GetAllProducts() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *productService) SearchProductsByName(name string) ([]Product, error) {
	return s.repo.SearchByName(name)
}

// HTTP Handlers

type productHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *productHandler {
	return &productHandler{service: service}
}

// helper to write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// helper to parse product id from URL
func parseIDparam(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid 'id' parameter: %w", err)
	}
	return id, nil
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	created, err := h.service.CreateProduct(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

func (h *productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDparam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "error fetching product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, p)
}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDparam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var update Product
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// To handle partial updates, negative price is treated as 'not provided'
	if update.Price < 0 {
		update.Price = -1
	}

	updated, err := h.service.UpdateProduct(id, &update)
	if err != nil {
		http.Error(w, "error updating product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if updated == nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDparam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "error deleting product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"deleted": true, "id": id})
}

func (h *productHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	nameQuery := r.URL.Query().Get("name")
	var (
		products []Product
		err      error
	)
	if nameQuery != "" {
		products, err = h.service.SearchProductsByName(nameQuery)
	} else {
		products, err = h.service.GetAllProducts()
	}
	if err != nil {
		http.Error(w, "error fetching products: "+err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, products)
}

func main() {
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to open DB connection: %v\n", err)
		return
	}
	defer db.Close()

	// verify DB connection
	if err = db.Ping(); err != nil {
		fmt.Printf("Failed to ping DB: %v\n", err)
		return
	}

	repo := NewPostgresProductRepository(db)
	service := NewProductService(repo)
	handler := NewProductHandler(service)

	r := mux.NewRouter()

	// RESTful routes
	r.HandleFunc("/products", handler.ListProducts).Methods("GET")   // list or search query param "name"
	r.HandleFunc("/products", handler.CreateProduct).Methods("POST") // create
	r.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
	r.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")

	fmt.Println("Product API server running at :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
