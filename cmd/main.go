package main

import (
	"log"
	"net"

	config "github.com/liju-github/EcommerceUserService/configs"
	"github.com/liju-github/EcommerceUserService/db"
	user "github.com/liju-github/CentralisedFoodbuddyMicroserviceProto/User"
	"github.com/liju-github/EcommerceUserService/repository"
	"github.com/liju-github/EcommerceUserService/service"
	util "github.com/liju-github/EcommerceUserService/utils"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	util.SetJWTSecretKey(cfg.JWTSecretKey)

	// Initialize database connection
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close(dbConn)

	// Initialize repository and service
	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userService)

	log.Println("User Service is running on gRPC port: " + cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server startup failed: %v", err)
	}
}
