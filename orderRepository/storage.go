package orderRepository

import (
	"context"
)

type Repository interface {
	CreateOrder(ctx context.Context, id string, info string) error
	FindAllOrders(ctx context.Context) (u []Order, err error)
	FindOrder(ctx context.Context, id string) (Order, error)
}
