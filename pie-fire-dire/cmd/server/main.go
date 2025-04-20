package main

import (
	"context"
	"log"
	grpc "newnok6/logic-test/pie-fire-dire/internal/adapter/handler/grpc"
	rest "newnok6/logic-test/pie-fire-dire/internal/adapter/handler/rest"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	grpcServer := grpc.NewGRPC("localhost:50051")
	httpServer := rest.NewHttpRest("localhost:8080")

	wg.Add(2)

	go func() {
		// Start gRPC server
		defer wg.Done()
		grpcServer.Start()
	}()

	go func() {
		// Start HTTP server
		defer wg.Done()
		httpServer.Start()
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	grpcServer.Stop()
	httpServer.Stop()

	wg.Wait()
	log.Println("Servers gracefully stopped.")
}
