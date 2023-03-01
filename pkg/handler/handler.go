package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/BisonVissarion/L0/orderRepository"
	"github.com/BisonVissarion/L0/pkg/cacheMemory"
	"github.com/BisonVissarion/L0/pkg/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repositry orderRepository.Repository
}

func (h *Handler) OrderHandler(c *cacheMemory.Cache) *gin.Engine {
	tablHtml, err := os.ReadFile("table.html")
	if err != nil {
		log.Fatal(err)
	}
	var i = template.Must(template.New("index").Parse(string(tablHtml)))

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("id")

		item, found := c.Get(id)
		info := fmt.Sprintf("%v", item)

		ord := orderRepository.Order{Info: info, ID: id}

		if !found {
			ord, err = h.Repositry.FindOrder(context.Background(), id)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "FindOrder failed: %v\n", err)
			} else {
				c.Set(ord.ID, ord.Info, 10*time.Minute)
			}
		}
		var modelOrd model.Order
		modelOrd.Items = make([]model.Items, 0)

		err := json.Unmarshal([]byte(ord.Info), &modelOrd)
		if err != nil {
			fmt.Fprintf(ctx.Writer, "unmarshal error: %s\n", err)
		}
		if err := i.Execute(ctx.Writer, modelOrd); err != nil {
			fmt.Printf("Execute error: %s\n", err)
		}
	})
	r.Run(":8080")

	return r
}
