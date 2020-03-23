package model

import (
	"packaging-service/constant"
)

type Order struct {
	ID              string
	Item            string
	Quantity        int
	PackagingStatus int
}

func NewOrder() *Order {
	return &Order{
		PackagingStatus: constant.UNPREPARED,
	}
}
