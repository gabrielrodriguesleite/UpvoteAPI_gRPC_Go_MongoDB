package main

import (
	"context"
	"io"
	"log"
	pb "upvote/grpc"
)

func doUpvote(c pb.UpvoteServiceClient) {
	log.Printf("Sending Upvote through gRPC...")

	req := &pb.UpvoteRequest{
		Vote: &pb.Vote{
			Email: "k.cogabriel@gmail.com",
			Score: 10,
		},
	}

	stream, err := c.Upvote(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling Upvote gRPC: %v", err)
	}

	for { // while true ?
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error while reading stream: %v", err)
		}

		log.Printf("Response from Upvote: %v\n", res.GetVote())
	}

}
