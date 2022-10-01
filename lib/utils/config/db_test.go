package config

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestConnDB(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name string
		args args
		want *mongo.Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnDB(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
