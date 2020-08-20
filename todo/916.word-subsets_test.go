Generated Test_wordSubsets
package main

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		A []string
		B []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSubsets(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
