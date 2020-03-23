package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	shipping "shipping-service/proto/shipping"
)

type Shipping struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Shipping) Call(ctx context.Context, req *shipping.Request, rsp *shipping.Response) error {
	log.Info("Received Shipping.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Shipping) Stream(ctx context.Context, req *shipping.StreamingRequest, stream shipping.Shipping_StreamStream) error {
	log.Infof("Received Shipping.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&shipping.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Shipping) PingPong(ctx context.Context, stream shipping.Shipping_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&shipping.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
