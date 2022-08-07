package main

import (
	"testing"
	pb "upvote/grpc"
)

func Test_server_Upvote(t *testing.T) {
	type fields struct {
		UnimplementedUpvoteServiceServer pb.UnimplementedUpvoteServiceServer
	}
	type args struct {
		req    *pb.UpvoteRequest
		server pb.UpvoteService_UpvoteServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedUpvoteServiceServer: tt.fields.UnimplementedUpvoteServiceServer,
			}
			if err := s.Upvote(tt.args.req, tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("server.Upvote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
