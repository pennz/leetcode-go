Generated Test_printVertically
package main

import (
	"reflect"
	"testing"
)

func Test_printVertically(t *testing.T) {
	type args struct {
		s string
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
			if got := printVertically(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
