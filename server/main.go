package main

import (
	"log"
	"net"

	pb "upvote/grpc"

	"google.golang.org/grpc"
)

func main() {
	// fmt.Println("server on...")

	// err := db.Create("k.cogabriel@gmail.com", 100)
	// if err != nil {
	// 	fmt.Println("um erro ocorreu:", err)
	// }

	// votes, err := db.Get()
	// if err != nil {
	// 	fmt.Println("um erro ocorreu:", err)
	// }

	// fmt.Println(votes)

	// fmt.Println("server off.")

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Fail to listen to: %v", err)
	}

	// grpc server
	s := grpc.NewServer()
	pb.RegisterUpvoteServiceServer(s, &server{})

	log.Println("Starting server, listen to port 50051.")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}
}
