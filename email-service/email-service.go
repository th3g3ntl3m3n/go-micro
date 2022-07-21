package main

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	emailservice "github.com/th3g3ntl3m3n/go-micro/email-service/proto"
)

type EmailService struct {
	emailRepo Respository
}

// Return a new handler
func NewEmailHandler(emailRepo Repository) *EmailService {
	return &EmailService{emailRepo}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *EmailService) Call(ctx context.Context, req *emailservice.Request, rsp *emailservice.Response) error {
	log.Info("Received EmailService.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *EmailService) Stream(ctx context.Context, req *emailservice.StreamingRequest, stream emailservice.EmailService_StreamStream) error {
	log.Infof("Received EmailService.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&emailservice.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *EmailService) PingPong(ctx context.Context, stream emailservice.EmailService_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&emailservice.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
