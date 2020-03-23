package main

import (
	"log"

	"packaging-service/handler"
	packaging "packaging-service/proto/packaging"
	"packaging-service/service"
	"packaging-service/subscriber"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

func main() {
	// New Service
	mservice := micro.NewService(
		micro.Name("go.micro.service.packaging"),
		micro.Version("latest"),
	)

	// Initialise service
	mservice.Init()

	// Register Handler
	packaging.RegisterPackagingHandler(mservice.Server(), new(handler.Packaging))

	// Create publisher
	pub := micro.NewPublisher("packaged-order", mservice.Client())

	// Register Struct as Subscriber
	packageService := service.CreatePackageService()
	sub := &subscriber.Packaging{
		PackageService: packageService,
		Publisher:      pub,
	}
	micro.RegisterSubscriber("new-order", mservice.Server(), sub, server.SubscriberQueue("new_order_queue"))

	// Run service
	if err := mservice.Run(); err != nil {
		log.Fatal(err)
	}
}
