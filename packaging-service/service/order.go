package service

import (
	"context"
	"encoding/json"
	"fmt"
	"packaging-service/constant"
	"packaging-service/model"

	packagingproto "packaging-service/proto/packaging"

	"github.com/go-log/log"
	"github.com/micro/go-micro/v2"
)

type PackageAction interface {
	ProcessPackage(publisher micro.Publisher, say string) error
}

type PackageService struct{}

func CreatePackageService() *PackageService {
	return &PackageService{}
}

func (ps *PackageService) ProcessPackage(publisher micro.Publisher, say string) (err error) {
	fmt.Println("Process order ....")
	var byteOrder []byte

	// Decode
	data := model.NewOrder()
	err = json.Unmarshal([]byte(say), data)
	if err != nil {
		return err
	}

	// Packaging
	data.PackagingStatus = constant.PACKAGED

	// Publish to packaging service
	byteOrder, err = json.Marshal(data)
	if err != nil {
		return
	}
	ev := &packagingproto.Message{
		Say: string(byteOrder),
	}
	if err := publisher.Publish(context.Background(), ev); err != nil {
		log.Logf("error publishing: %v", err)
	}
	return
}
