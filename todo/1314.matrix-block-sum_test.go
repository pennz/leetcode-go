package main

import (
	"reflect"
	"testing"
)

func Test_matrixBlockSum(t *testing.T) {
	type args struct {
		mat [][]int
		K   int
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
			if got := matrixBlockSum(tt.args.mat, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
