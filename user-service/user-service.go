package main

import (
	"context"

	pb "github.com/th3g3ntl3m3n/go-micro/user-service/proto"
)

type UserService struct {
	userRepo Repository
}

// Return a new handler
func NewUserHandler(userRepo Repository) *UserService {
	return &UserService{userRepo}
}

func (e *UserService) Get(context.Context, *pb.GetRequest, *pb.UserResponse) error {

	return nil
}
