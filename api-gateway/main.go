package main

import (
	"context"
	"fmt"

	api "api-gateway/proto/api"

	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.api"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// // Register Handler
	// api.RegisterApiHandler(service.Server(), new(handler.Api))

	// // Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.service.api", service.Server(), new(subscriber.Api))

	// Create Order Service client
	orderService := api.NewOrderService("go.micro.service.order", service.Client())

	// Call Order Service
	resp, err := orderService.Call(context.TODO(), &api.Request{Name: "Bill"})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resp)

	var orderResp *api.OrderResponse
	orderResp, err = orderService.AddOrder(context.TODO(), &api.OrderRequest{Item: "Meaty Mix", Quantity: 100})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(orderResp)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
