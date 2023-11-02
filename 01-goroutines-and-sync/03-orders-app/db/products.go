package db

import (
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/kxplxn/go-concurrency/01-goroutines-and-sync/03-orders-app/models"
	"github.com/kxplxn/go-concurrency/01-goroutines-and-sync/03-orders-app/utils"
)

type ProductDB struct {
	products sync.Map
}

// NewProducts creates a new empty product DB
func NewProducts() (*ProductDB, error) {
	p := &ProductDB{}
	// load start position
	if err := utils.ImportProducts(&p.products); err != nil {
		return nil, err
	}

	return p, nil
}

// Exists checks whether a product with a give id exists
func (p *ProductDB) Exists(id string) error {
	if _, ok := p.products.Load(id); !ok {
		return fmt.Errorf("no product found for id %s", id)
	}

	return nil
}

// Find returns a given product if one exists
func (p *ProductDB) Find(id string) (models.Product, error) {
	pp, ok := p.products.Load(id)
	if !ok {
		return models.Product{}, fmt.Errorf("no product found for id %s", id)
	}

	return toProduct(pp), nil
}

// Upsert creates or updates a product in the orders DB
func (p *ProductDB) Upsert(prod models.Product) {
	p.products.Store(prod.ID, prod)
}

// FindAll returns all products in the system
func (p *ProductDB) FindAll() []models.Product {
	var allProducts []models.Product
	p.products.Range(func(key, value any) bool {
		allProducts = append(allProducts, toProduct(value))
		return true
	})
	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].ID < allProducts[j].ID
	})
	return allProducts
}

func toProduct(pp any) models.Product {
	prod, ok := pp.(models.Product)
	if !ok {
		log.Fatalf("error casting %v to product", pp)
	}
	return prod
}
