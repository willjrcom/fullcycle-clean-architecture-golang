package repository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/willjrcom/fullcycle-clean-architecture/golang/internal/domain"
)

type OrderRepositoryBun struct {
	mu sync.Mutex
	db *bun.DB
}

func NewOrderRepositoryBun(db *bun.DB) *OrderRepositoryBun {
	return &OrderRepositoryBun{db: db}
}

func (r *OrderRepositoryBun) NewOrder(ctx context.Context, order *domain.Order) (uuid.UUID, error) {
	r.mu.TryLock()
	defer r.mu.Unlock()

	if _, err := r.db.NewInsert().Model(order).Exec(ctx); err != nil {
		return uuid.Nil, err
	}

	return order.ID, nil
}

func (r *OrderRepositoryBun) GetAllOrders(ctx context.Context) ([]domain.Order, error) {
	r.mu.TryLock()
	defer r.mu.Unlock()

	orders := []domain.Order{}

	if err := r.db.NewSelect().Model(&orders).Scan(ctx); err != nil {
		return nil, err
	}

	return orders, nil
}
