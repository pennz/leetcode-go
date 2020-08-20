Generated Test_movesToStamp
package main

import (
	"reflect"
	"testing"
)

func Test_movesToStamp(t *testing.T) {
	type args struct {
		stamp  string
		target string
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
			if got := movesToStamp(tt.args.stamp, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("movesToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
