package main

import (
	"reflect"
	"testing"
)

func Test_pathsWithMaxScore(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathsWithMaxScore(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathsWithMaxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
