package main

import (
	"testing"
	pb "upvote/grpc"
)

func Test_doUpvote(t *testing.T) {
	type args struct {
		c pb.UpvoteServiceClient
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doUpvote(tt.args.c)
		})
	}
}
