package orderservice

import (
	"context"
	"fmt"
	"order-service/tool"

	proto "order-service/proto"

	log "github.com/sirupsen/logrus"
)

func Run() (err error) {
	// Connect to database
	dbConnector := tool.NewDBConnector()
	db, err := tool.ConnectDB(dbConnector)
	if err != nil {
		log.WithError(err).Fatal("Cannot connect to database")
		return
	}

	// Create a server
	proto.RegisterOrderServer(service.Server(), new(Greeter))

	service.Init()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func (g *Greeter) AddOrder(context.Context, req *OrderRequest) (rsp *OrderResponse, err error)
	rsp.Id = "id"
	rsp.Item = req.Item
	rsp.Quantity = req.Quantity

	return nil
}

type Greeter struct {
}
