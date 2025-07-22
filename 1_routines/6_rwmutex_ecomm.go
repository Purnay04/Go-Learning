package main

import (
	"fmt"
	"sync"
	"time"
)

type Product struct {
	ID   int
	Name string
}

type Catalog struct {
	products map[int]Product
	mu       sync.RWMutex
}

// GetProduct safely allows multiple readers
func (c *Catalog) GetProduct(id int) (Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, ok := c.products[id]
	return p, ok
}

// AddOrUpdateProduct grants exclusive write access
func (c *Catalog) AddOrUpdateProduct(p Product) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products[p.ID] = p
}

func main() {
	catalog := &Catalog{
		products: make(map[int]Product),
	}

	// Add some products (write lock)
	catalog.AddOrUpdateProduct(Product{ID: 1, Name: "Laptop"})
	catalog.AddOrUpdateProduct(Product{ID: 2, Name: "Phone"})

	var wg sync.WaitGroup

	// Simulate many users reading concurrently (RLock)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if p, ok := catalog.GetProduct(1); ok {
				fmt.Printf("User %d sees product: %s\n", id, p.Name)
			}
		}(i)
	}

	// Simulate an admin updating a product (Lock)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		catalog.AddOrUpdateProduct(Product{ID: 1, Name: "New Laptop"})
		fmt.Println("Admin updated Laptop to 'New Laptop'")
	}()

	wg.Wait()
}
