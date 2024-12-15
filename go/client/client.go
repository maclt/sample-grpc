package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "maclt/sample-grpc/location"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the RouteGuide service
	client := pb.NewLocationClient(conn)

	// Call the GetFeature method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	point := &pb.Point{Lat: 35.65861098290943, Lng: 139.74543289643074}
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("Could not get feature: %v", err)
	}

	log.Printf("Feature received: %s", feature.Name+" in "+feature.Country)
}
