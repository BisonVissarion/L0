package order

import (
	"context"
	"fmt"

	"github.com/BisonVissarion/L0/orderRepository"
	"github.com/BisonVissarion/L0/pkg/clientPostgres"
	"github.com/jackc/pgconn"
)

type repository struct {
	client clientPostgres.Client
}

func NewRepository(client clientPostgres.Client) orderRepository.Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) FindOrder(ctx context.Context, id string) (orderRepository.Order, error) {
	q := `SELECT id, info 
	      FROM orders 
		  WHERE id = $1`
	var ath orderRepository.Order
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Info)
	if err != nil {
		return orderRepository.Order{}, err
	}
	return ath, nil
}

func (r *repository) FindAllOrders(ctx context.Context) (u []orderRepository.Order, err error) {
	q := `SELECT id, info 
	      FROM orders;`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	orders := make([]orderRepository.Order, 0)
	for rows.Next() {
		var ath orderRepository.Order
		err = rows.Scan(&ath.ID, &ath.Info)
		if err != nil {
			return nil, err
		}
		orders = append(orders, ath)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *repository) CreateOrder(ctx context.Context, id string, info string) error {
	q := `INSERT INTO orders (id, info) 
	      VALUES ($1, $2) 
		  RETURNING id`
	err := r.client.QueryRow(ctx, q, id, info).Scan()
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}
	return nil
}
