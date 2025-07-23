package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// add a function to search the product by its name using a like operator

// disk --> 'pen' or a 'pencil'

// Product represents a product entity
type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

// ProductRepository defines DB methods for products
type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(p *Product) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO products (name, price, category) VALUES ($1, $2, $3) RETURNING product_id",
		p.Name, p.Price, p.Category).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductRepository) GetByID(id int) (*Product, error) {
	p := &Product{}
	row := r.db.QueryRow("SELECT product_id, name, price, category FROM products WHERE product_id = $1", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Category)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) UpdatePrice(id int, price float64) error {
	_, err := r.db.Exec("UPDATE products SET price = $1 WHERE product_id = $2", price, id)
	return err
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	rows, err := r.db.Query("SELECT product_id, name, price, category FROM products")
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

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE product_id = $1", id)
	return err
}

// ProductService provides business logic on products
type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float64, category string) (*Product, error) {
	p := &Product{
		Name:     name,
		Price:    price,
		Category: category,
	}
	id, err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}
	p.ID = id
	return p, nil
}

func (s *ProductService) GetProduct(id int) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProductPrice(id int, price float64) error {
	return s.repo.UpdatePrice(id, price)
}

func (s *ProductService) ListProducts() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}

// main function initializes DB, service, and demonstrates usage
func main() {
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	service := NewProductService(repo)

	// Create a new product
	product, err := service.CreateProduct("Book", 299.99, "Books")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted Product: %+v\n", product)

	// Read the product by ID
	pFromDB, err := service.GetProduct(product.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read Product: %+v\n", pFromDB)

	// Update product price
	err = service.UpdateProductPrice(product.ID, 249.99)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated product price.")

	// List all products
	products, err := service.ListProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All products:")
	for _, prod := range products {
		fmt.Printf("%+v\n", prod)
	}

	// Delete the product
	err = service.DeleteProduct(product.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted product.")
}

// main --> Service --> Repository/DAO --> execs to the DB
