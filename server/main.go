package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "upvote/grpc"

	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
)

var API_HOST = ""
var API_PORT = ""
var API_URI = ""

func init() {
	gotenv.Load()
	API_HOST = os.Getenv("API_HOST")
	API_PORT = os.Getenv("API_PORT")

	if API_HOST == "" {
		API_HOST = "localhost"
		API_PORT = "50051"
	}

	API_URI = fmt.Sprintf("%s:%s", API_HOST, API_PORT)
}

func main() {

	lis, err := net.Listen("tcp", API_URI)
	if err != nil {
		log.Fatalf("Fail to listen to: %v", err)
	}

	// grpc server
	s := grpc.NewServer()
	pb.RegisterUpvoteServiceServer(s, &server{})

	log.Printf("Starting server, listen to port %s.", API_PORT)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}
}
