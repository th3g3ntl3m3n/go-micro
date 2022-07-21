package main

import (
	"context"

	pb "github.com/th3g3ntl3m3n/go-micro/email-service/proto"
)

type EmailService struct {
	emailRepo Repository
}

// Return a new handler
func NewEmailHandler(emailRepo Repository) *EmailService {
	return &EmailService{emailRepo}
}

func (e *EmailService) Get(ctx context.Context, in *pb.UserRequest, rsp *pb.EmailResponse) error {
	rsp.Email = e.emailRepo.GetByUserID(in.Id)
	return nil
}
