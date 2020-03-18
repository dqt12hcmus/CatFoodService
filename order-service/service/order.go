package service

import (
	"fmt"
	"order-service/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderAction interface {
	ProcessOrder(db *mongo.Database, order *model.Order)
}

type OrderService struct{}

func CreateOrderService() *OrderService {
	return &OrderService{}
}

func (os *OrderService) ProcessOrder(db *mongo.Database, order *model.Order) {
	fmt.Println("Process order ....")
	return
}
