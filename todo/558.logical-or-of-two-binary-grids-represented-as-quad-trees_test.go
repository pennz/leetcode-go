package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		quadTree1 *Node
		quadTree2 *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.quadTree1, tt.args.quadTree2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
