package main

import (
	"packaging-service/handler"
	"packaging-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	packaging "packaging-service/proto/packaging"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.packaging"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	packaging.RegisterPackagingHandler(service.Server(), new(handler.Packaging))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("new-order", service.Server(), new(subscriber.Packaging), server.SubscriberQueue("new_order_queue"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
