package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	packaging "packaging-service/proto/packaging"
	"packaging-service/service"
)

type Packaging struct {
	PackageService *service.PackageService
	Publisher      micro.Publisher
}

func (e *Packaging) Handle(ctx context.Context, msg *packaging.Message) error {
	e.PackageService.ProcessPackage(e.Publisher, msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *packaging.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
