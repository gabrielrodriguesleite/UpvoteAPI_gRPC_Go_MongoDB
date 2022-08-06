package main

import (
	"fmt"
	"upvote/db"
	pb "upvote/grpc"
)

type server struct {
	pb.UnimplementedUpvoteServiceServer
}

func (*server) Upvote(req *pb.UpvoteRequest, server pb.UpvoteService_UpvoteServer) error {

	fmt.Printf("Upvote was invoked with: %v\n", req)

	// votar
	err := db.Create(req.GetVote().GetEmail(), int(req.GetVote().GetScore()))
	if err != nil {
		fmt.Println("Error at Create: ", err)
	}

	// responder
	votes, err := db.Get()
	if err != nil {
		fmt.Println("Error at Get: ", err)
	}

	// non concurrent stream
	for _, x := range votes {
		resp := pb.UpvoteResponse{
			Vote: &pb.Vote{
				Email: x.Email, Score: int32(x.Score),
			},
		}
		server.Send(&resp)
	}

	return err
}
