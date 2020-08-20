package main

import (
	"reflect"
	"testing"
)

func Test_queensAttacktheKing(t *testing.T) {
	type args struct {
		queens [][]int
		king   []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queensAttacktheKing(tt.args.queens, tt.args.king); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queensAttacktheKing() = %v, want %v", got, tt.want)
			}
		})
	}
}
