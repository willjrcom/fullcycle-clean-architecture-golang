package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/willjrcom/fullcycle-clean-architecture/golang/internal/domain"
)

type Service struct {
	r domain.Repository
}

func NewService(r domain.Repository) *Service {
	return &Service{r: r}
}

func (s *Service) NewOrder(ctx context.Context, orderCommonAttributes domain.OrderCommonAttributes) (uuid.UUID, error) {
	order := domain.NewOrder(orderCommonAttributes)
	return s.r.NewOrder(ctx, order)
}

func (s *Service) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return s.r.GetAllOrders(ctx)
}
