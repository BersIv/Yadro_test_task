package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "yadro_test_task/gen/go/hostconfig"
	"yadro_test_task/internal/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcHostPort = "0.0.0.0:8081"
)

func main() {
	grpcServer := grpc.NewServer()
	lisen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalf("Failed listen: %v", err)
	}

	pb.RegisterHostConfigServer(grpcServer, &service.Server{})

	go func() {
		log.Printf("Starting gRPC server :%v", grpcHostPort)
		if err := grpcServer.Serve(lisen); err != nil {
			log.Fatalf("Failed serve: %v", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterHostConfigHandlerFromEndpoint(ctx, mux, grpcHostPort, opts)
	if err != nil {
		log.Fatalf("Failed register: %v", err)
	}
	log.Printf("Starting REST server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed listen and serve: %v", err)
	}

}
