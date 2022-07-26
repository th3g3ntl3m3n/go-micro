package main

import (
	emailPb "github.com/th3g3ntl3m3n/go-micro/email-service/proto"
	pb "github.com/th3g3ntl3m3n/go-micro/user-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user-service"),
	)

	emailClient := emailPb.NewEmailService("email-service", srv.Client())

	db := make(UsersDB)
	db["1"] = &User{ID: "1", Name: "Vikas"}
	db["2"] = &User{ID: "2", Name: "Test 2"}

	userRepo := New(db)

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), NewUserHandler(userRepo, emailClient))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
