package stats

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kxplxn/go-concurrency/03-concurrency-patterns/02-worker-pools/models"
)

const WorkerCount = 3

type statsService struct {
	result    Result
	processed <-chan models.Order
	done      <-chan struct{}
	pStats    chan models.Statistics
}

type StatsService interface {
	GetStats() models.Statistics
}

func New(processed <-chan models.Order, done <-chan struct{}) StatsService {
	s := statsService{
		result:    &result{},
		processed: processed,
		done:      done,
		pStats:    make(chan models.Statistics, WorkerCount),
	}
	for i := 0; i < WorkerCount; i++ {
		go s.processStats()
	}
	go s.reconcile()
	return &s
}

// processStats is the overall processing method that listens to incoming orders
func (s *statsService) processStats() {
	fmt.Println("Stats processing started!")
	for {
		select {
		case order := <-s.processed:
			pstats := s.processOrder(order)
			s.pStats <- pstats
		case <-s.done:
			fmt.Println("Stats processing stopped!")
			return
		}
	}
}

// reconcile is a helper method which saves stats object
// back into the statisticsService
func (s *statsService) reconcile() {
	fmt.Println("Reconcile started!")
	for {
		select {
		case p := <-s.pStats:
			s.result.Combine(p)
		case <-s.done:
			fmt.Println("Reconcile stopped!")
			return
		}
	}
}

// processOrder is a helper method that incorporates the current order in the stats service
func (s *statsService) processOrder(order models.Order) models.Statistics {
	// simulate processing as a costly operation
	randomSleep()
	// completed orders increment add to the revenue
	if order.Status == models.OrderStatus_Completed {
		return models.Statistics{
			CompletedOrders: 1,
			Revenue:         *order.Total,
		}
	}
	// otherwise the order is rejected
	return models.Statistics{
		RejectedOrders: 1,
	}
}

// GetStats returns the latest order stats
func (s *statsService) GetStats() models.Statistics {
	return s.result.Get()
}

func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}
