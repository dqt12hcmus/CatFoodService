package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	packaging "packaging-service/proto/packaging"
)

type Packaging struct{}

func (e *Packaging) Handle(ctx context.Context, msg *packaging.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *packaging.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
