package pb

import (
	"context"

	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/domain"
	"github.com/willjrcom/fullcycle-clean-architecture-golang/internal/usecase"
)

type ServiceGrpc struct {
	UnimplementedOrderServiceServer
	s *usecase.Service
}

func NewServiceGrpc(s *usecase.Service) *ServiceGrpc {
	return &ServiceGrpc{s: s}
}

func (s *ServiceGrpc) NewOrder(ctx context.Context, input *CreateOrderRequest) (*OrderResponse, error) {
	orderCommonAttributes := domain.OrderCommonAttributes{
		Name:  input.Name,
		Total: float64(input.Total),
	}

	id, error := s.s.NewOrder(ctx, orderCommonAttributes)

	if error != nil {
		return nil, error
	}

	orderResponse := &OrderResponse{
		Order: &Order{
			Id:    id.String(),
			Name:  input.Name,
			Total: input.Total,
		},
	}
	return orderResponse, nil
}

func (s *ServiceGrpc) ListOrders(ctx context.Context, _ *Blank) (*OrderListResponse, error) {
	orders, err := s.s.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	pbOrders := []*Order{}

	for _, order := range orders {
		pbOrders = append(pbOrders, &Order{
			Id:    order.ID.String(),
			Name:  order.Name,
			Total: float32(order.Total),
		})
	}

	orderListResponse := &OrderListResponse{
		Orders: pbOrders,
	}
	return orderListResponse, nil
}
