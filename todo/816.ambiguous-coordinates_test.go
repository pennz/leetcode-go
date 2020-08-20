Generated Test_ambiguousCoordinates
package main

import (
	"reflect"
	"testing"
)

func Test_ambiguousCoordinates(t *testing.T) {
	type args struct {
		S string
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
			if got := ambiguousCoordinates(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ambiguousCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
