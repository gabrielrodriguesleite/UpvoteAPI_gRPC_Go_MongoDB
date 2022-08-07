package db

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		collection string
		data       any
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
			if err := Insert(tt.args.collection, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name     string
		args     args
		wantList []Vote
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := List(tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("List() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
