package main

import (
	"context"
	"order-service/handler"
	"order-service/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"order-service/config"
	order "order-service/proto/order"
	"order-service/service"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.DB.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(config.DB.DBName)

	// New Service
	mservice := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
	)

	// Initialize service
	mservice.Init()

	// Create publisher
	pub := micro.NewPublisher("new-order", mservice.Client())

	// Register Handler
	orderService := &service.OrderService{}
	handler := &handler.Order{
		OrderService: orderService,
		DB:           database,
		Publisher:    pub,
	}
	order.RegisterOrderHandler(mservice.Server(), handler)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.order", mservice.Server(), new(subscriber.Order))

	// Run service
	if err := mservice.Run(); err != nil {
		log.Fatal(err)
	}
}
