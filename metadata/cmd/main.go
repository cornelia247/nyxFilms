package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/metadata/internal/controller/metadata"
	grpchandler "github.com/cornelia247/nyxfilms/metadata/internal/handler/grpc"
	"github.com/cornelia247/nyxfilms/metadata/internal/repository/memory"
	"github.com/cornelia247/nyxfilms/pkg/discovery"
	"github.com/cornelia247/nyxfilms/pkg/discovery/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting the metadata service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	ctrl := metadata.New(repo)

	// HTTP JSON HANDLER.
	// h := httphandler.New(ctrl)
	// http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	// if err := http.ListenAndServe(fmt.Sprintf(":%d",port), nil); err != nil {
	// 	panic(err)
	// }

	// PROTOBUFF HANDLER.
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterMetadataServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}

}
