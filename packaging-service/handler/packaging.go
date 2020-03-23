package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	packaging "packaging-service/proto/packaging"
)

type Packaging struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Packaging) Call(ctx context.Context, req *packaging.Request, rsp *packaging.Response) error {
	log.Info("Received Packaging.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Packaging) Stream(ctx context.Context, req *packaging.StreamingRequest, stream packaging.Packaging_StreamStream) error {
	log.Infof("Received Packaging.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&packaging.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Packaging) PingPong(ctx context.Context, stream packaging.Packaging_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&packaging.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
