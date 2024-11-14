package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/cornelia247/nyxfilms/gen"
	"github.com/cornelia247/nyxfilms/pkg/discovery"
	"github.com/cornelia247/nyxfilms/pkg/discovery/consul"
	"github.com/cornelia247/nyxfilms/rating/internal/controller/rating"
	grpchandler "github.com/cornelia247/nyxfilms/rating/internal/handler/grpc"
	"github.com/cornelia247/nyxfilms/rating/internal/repository/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
)

const serviceName = "rating"

func main() {

	f, err := os.Open("base.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var cfg config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	port := cfg.API.port
	log.Printf("Starting he rating service on port %d", port)
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

	repo, err := mysql.New()
	if err != nil {
		panic(err)
	}
	ctrl := rating.New(repo, nil)

	// HTTP HANDLER.
	// h := httphandler.New(ctrl)
	// http.Handle("/rating", http.HandlerFunc(h.Handle))
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
	// 	panic(err)
	// }

	// PROTO BUFF HANDLER.
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterRatingServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
