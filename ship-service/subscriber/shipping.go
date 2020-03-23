package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	shipping "shipping-service/proto/shipping"
)

type Shipping struct{}

func (e *Shipping) Handle(ctx context.Context, msg *shipping.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	log.Info("Shipping packaged order")
	return nil
}

func Handler(ctx context.Context, msg *shipping.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
