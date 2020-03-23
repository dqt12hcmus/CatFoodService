package main

import (
	"packaging-service/handler"
	"packaging-service/service"
	"packaging-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"

	packaging "packaging-service/proto/packaging"
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
	sub := subscriber.Packaging{
		PackageService: packageService,
		Publisher:      pub,
	}
	micro.RegisterSubscriber("packaged-order", mservice.Server(), sub, server.SubscriberQueue("packaged_order_queue"))

	// Run service
	if err := mservice.Run(); err != nil {
		log.Fatal(err)
	}
}
