package domain

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	NewOrder(ctx context.Context, order *Order) (uuid uuid.UUID, err error)
	GetAllOrders(ctx context.Context) (orders []Order, err error)
	GetByID(ctx context.Context, id string) (order *Order, err error)
}
