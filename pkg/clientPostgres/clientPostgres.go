package clientPostgres

import (
	"context"
	"fmt"
	"os"

	"github.com/BisonVissarion/L0/orderRepository/config"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func NewClient(sc config.StorageConfig) (c *pgx.Conn) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
