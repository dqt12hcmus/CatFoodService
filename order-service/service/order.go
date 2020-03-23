package service

import (
	"context"
	"encoding/json"
	"fmt"
	"order-service/config"
	"order-service/model"

	orderproto "order-service/proto/order"

	"github.com/go-log/log"
	"github.com/micro/go-micro/v2"
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

func (os *OrderService) ProcessOrder(publisher micro.Publisher, db *mongo.Database, order *model.Order) (err error) {
	fmt.Println("Process order ....")
	var result *mongo.InsertOneResult
	var byteOrder []byte

	// Save to database
	collection := db.Collection(config.DB.Collection)
	result, err = collection.InsertOne(context.TODO(), order)
	if err != nil {
		return
	}
	objectID := result.InsertedID.(primitive.ObjectID)
	order.ID = objectID.String()

	// Publish to packaging service
	byteOrder, err = json.Marshal(order)
	if err != nil {
		return
	}
	ev := &orderproto.Message{
		Say: string(byteOrder),
	}
	if err := publisher.Publish(context.Background(), ev); err != nil {
		log.Logf("error publishing: %v", err)
	}
	return
}
