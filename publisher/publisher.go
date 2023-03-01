package main

import (
	"fmt"
	"log"
	"os"

	stan "github.com/nats-io/stan.go"
)

func main() {
	stanConn, _ := stan.Connect("test-cluster", "2")
	defer stanConn.Close()
	for i := 0; i < 2; i++ {
		dataOrders, err := os.ReadFile(fmt.Sprintf("file%b.json", i))
		if err != nil {
			log.Fatal(err)
		}
		stanConn.Publish("foo", dataOrders)
	}
}
