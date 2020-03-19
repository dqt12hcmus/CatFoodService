package service

import (
	"context"
	"fmt"
	"order-service/config"
	"order-service/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderAction interface {
	ProcessOrder(db *mongo.Database, order *model.Order) error
}

type OrderService struct{}

func CreateOrderService() *OrderService {
	return &OrderService{}
}

func (os *OrderService) ProcessOrder(db *mongo.Database, order *model.Order) (err error) {
	fmt.Println("Process order ....")
	var result *mongo.InsertOneResult

	collection := db.Collection(config.DB.Collection)
	result, err = collection.InsertOne(context.TODO(), order)
	if err != nil {
		return
	}
	objectID := result.InsertedID.(primitive.ObjectID)
	order.ID = objectID.String()
	return
}
