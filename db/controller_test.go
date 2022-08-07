package db

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		email string
		score int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.email, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		wantVotes []Vote
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVotes, err := Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVotes, tt.wantVotes) {
				t.Errorf("Get() = %v, want %v", gotVotes, tt.wantVotes)
			}
		})
	}
}
