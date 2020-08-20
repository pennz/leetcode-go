package main

import (
	"reflect"
	"testing"
)

func Test_basicCalculatorIV(t *testing.T) {
	type args struct {
		expression string
		evalvars   []string
		evalints   []int
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
			if got := basicCalculatorIV(tt.args.expression, tt.args.evalvars, tt.args.evalints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("basicCalculatorIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
