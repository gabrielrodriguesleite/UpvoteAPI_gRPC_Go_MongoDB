package main

import (
	"fmt"
	"log"

	pb "upvote/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Starting Client...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could note connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewUpvoteServiceClient(cc)

	doUpvote(c)

}
