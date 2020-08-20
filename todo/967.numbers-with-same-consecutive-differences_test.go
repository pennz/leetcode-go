package main

import (
	"reflect"
	"testing"
)

func Test_numsSameConsecDiff(t *testing.T) {
	type args struct {
		N int
		K int
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
			if got := numsSameConsecDiff(tt.args.N, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsSameConsecDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
