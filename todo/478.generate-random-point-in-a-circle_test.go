Generated TestConstructor
Generated TestSolution_RandPoint
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		radius   float64
		x_center float64
		y_center float64
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.radius, tt.args.x_center, tt.args.y_center); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_RandPoint(t *testing.T) {
	tests := []struct {
		name string
		this *Solution
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.RandPoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution.RandPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
