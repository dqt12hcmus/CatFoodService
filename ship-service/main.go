package main

import (
	"shipping-service/handler"
	"shipping-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	shipping "shipping-service/proto/shipping"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.shipping"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	shipping.RegisterShippingHandler(service.Server(), new(handler.Shipping))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("packaged-order", service.Server(), new(subscriber.Shipping), server.SubscriberQueue("packaged_order_queue"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
