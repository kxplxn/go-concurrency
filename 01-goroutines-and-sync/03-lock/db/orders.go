package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/kxplxn/go-concurrency/01-goroutines-and-sync/03-lock/models"
)

type OrderDB struct{ placedOrders sync.Map }

// NewOrders creates a new empty order service
func NewOrders() *OrderDB { return &OrderDB{} }

// Find order for a given id, if one exists
func (o *OrderDB) Find(id string) (models.Order, error) {
	po, ok := o.placedOrders.Load(id)
	if !ok {
		return models.Order{}, fmt.Errorf("no order found for %s order id", id)
	}

	return toOrder(po), nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders.Store(order.ID, order)
}

func toOrder(po any) models.Order {
	order, ok := po.(models.Order)
	if !ok {
		log.Fatalf("error casting %v to order", order)
	}
	return order
}
