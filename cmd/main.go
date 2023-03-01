package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	stan "github.com/nats-io/stan.go"

	"github.com/BisonVissarion/L0/orderRepository/config"
	"github.com/BisonVissarion/L0/orderRepository/order"
	"github.com/BisonVissarion/L0/pkg/cacheMemory"
	"github.com/BisonVissarion/L0/pkg/clientPostgres"
	"github.com/BisonVissarion/L0/pkg/handler"
	"github.com/BisonVissarion/L0/pkg/model"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "1")

	conn := clientPostgres.NewClient(*config.GetConfig())

	var h handler.Handler
	h.Repositry = order.NewRepository(conn)

	orders, err := h.Repositry.FindAllOrders(context.Background())
	if err != nil {
		fmt.Printf("FindAllOrders failed: %v\n", err)
	}

	с := cacheMemory.New()

	for _, order := range orders {
		с.Set(order.ID, order.Info, 10*time.Minute)
	}

	sc.Subscribe("foo", func(m *stan.Msg) {
		var dataOrd model.Order
		err := json.Unmarshal(m.Data, &dataOrd)
		if err != nil {
			fmt.Printf("unmarshal error: %s\n", err)
		}
		info, _ := json.Marshal(dataOrd)
		h.Repositry.CreateOrder(context.Background(), dataOrd.Order_uid, string(info))
	})

	h.OrderHandler(с)
	sc.Close()
	conn.Close(context.Background())
}
