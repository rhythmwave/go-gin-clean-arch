package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"eira/document/config"
	"eira/document/router"
)

const (
	defaultTimeout = 30
	defaultPort    = ":8080"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	godotenv.Load()
	config.LoadConfig()
	// db := config.InitDB()
	var db *gorm.DB

	port := os.Getenv("SERVE_PORT")
	if port == "" {
		port = defaultPort
	}

	r := router.SetupRouter(db)
	// go func() {
	// 	if err := r.Run(port); err != nil {
	// 		log.Fatalf("Failed to run HTTP server: %v", err)
	// 	}
	// }()

	r.Run(port)
	// Setup gRPC server
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()
	// grpc.RegisterUserServer(grpcServer, userUseCase)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve gRPC server: %v", err)
	// }

}
