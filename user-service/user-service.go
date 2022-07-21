package main

import (
	"context"
	"fmt"

	emailPb "github.com/th3g3ntl3m3n/go-micro/email-service/proto"
	pb "github.com/th3g3ntl3m3n/go-micro/user-service/proto"
)

type UserService struct {
	userRepo    Repository
	emailClient emailPb.EmailService
}

// Return a new handler
func NewUserHandler(userRepo Repository, emailClient emailPb.EmailService) *UserService {
	return &UserService{userRepo, emailClient}
}

func (e *UserService) Get(ctx context.Context, in *pb.GetRequest, rsp *pb.UserResponse) error {
	user := e.userRepo.GetByID(in.Id)
	if user == nil {
		return fmt.Errorf("user not found")
	}

	resp, err := e.emailClient.Get(ctx, &emailPb.UserRequest{Id: user.ID})
	if err != nil {
		return fmt.Errorf("Error calling emailclient")
	}
	rsp.Id = user.ID
	rsp.Name = user.Name
	rsp.Email = resp.Email
	return nil
}
