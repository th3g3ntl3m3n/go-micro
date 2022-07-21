package main

import (
	pb "github.com/th3g3ntl3m3n/go-micro/user-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user-service"),
	)

	db := make(UsersDB)
	userRepo := New(db)

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), NewUserHandler(userRepo))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
