package main

import (
	"log"
	"net"

	user "github.com/liju-github/CentralisedFoodbuddyMicroserviceProto/User"
	config "github.com/liju-github/FoodBuddyMicroserviceUser/configs"
	"github.com/liju-github/FoodBuddyMicroserviceUser/db"
	"github.com/liju-github/FoodBuddyMicroserviceUser/repository"
	"github.com/liju-github/FoodBuddyMicroserviceUser/service"
	util "github.com/liju-github/FoodBuddyMicroserviceUser/utils"
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
	listener, err := net.Listen("tcp", ":"+cfg.USERGRPCPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userService)

	log.Println("User Service is running on gRPC port: " + cfg.USERGRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server startup failed: %v", err)
	}
}
