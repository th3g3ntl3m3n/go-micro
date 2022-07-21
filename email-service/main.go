package main

import (
	pb "github.com/th3g3ntl3m3n/go-micro/email-service/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("email-service"),
	)

	db := make(EmailsDB)
	db["1"] = "vikas@test.com"
	db["2"] = "vikas@url.com"
	emailRepo := NewEmailRepository(db)

	// Register handler
	pb.RegisterEmailServiceHandler(srv.Server(), NewEmailHandler(emailRepo))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
