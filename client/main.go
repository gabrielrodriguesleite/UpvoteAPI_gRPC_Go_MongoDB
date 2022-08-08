package main

import (
	"fmt"
	"log"
	"os"

	pb "upvote/grpc"

	"github.com/subosito/gotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	fmt.Println("Starting Client...")

	cc, err := grpc.Dial(API_URI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could note connect: %v", err)
	}
	defer cc.Close()

	c := pb.NewUpvoteServiceClient(cc)

	doUpvote(c)

}
