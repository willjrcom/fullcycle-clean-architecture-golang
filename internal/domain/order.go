package domain

import (
	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID `bun:"id,type:uuid,pk,notnull" json:"id"`
	OrderCommonAttributes
}

type OrderCommonAttributes struct {
	Name  string  `bun:"name,notnull" json:"name"`
	Total float64 `bun:"total,notnull" json:"total"`
}

func NewOrder(orderCommonAttributes OrderCommonAttributes) *Order {
	return &Order{
		ID:                    uuid.New(),
		OrderCommonAttributes: orderCommonAttributes,
	}
}
