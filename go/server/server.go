package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "maclt/sample-grpc/location"
)

// Server struct implements the LocationServer interface
type locationServer struct {
	pb.UnimplementedLocationServer
}

// Implement GetFeature RPC
func (s *locationServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Printf("Received request for Point: %v, %v", point.Lat, point.Lng)
	return &pb.Feature{Name: "Tokyo Tower", Country: "Japan"}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register the location server
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLocationServer(grpcServer, &locationServer{})
	log.Println("gRPC server listening on port 50051")

	// Serve gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
