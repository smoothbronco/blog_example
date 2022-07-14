package main

import (
	"log"
	"net"

	"github.com/smoothbronco/blog_example/article/pb"
	"github.com/smoothbronco/blog_example/article/repository"
	"github.com/smoothbronco/blog_example/article/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}

	service := service.NewService(repository)

	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	log.Println("Listening on port 8080...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
