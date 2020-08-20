Generated TestConstructor
Generated TestSubrectangleQueries_UpdateSubrectangle
Generated TestSubrectangleQueries_GetValue
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		rectangle [][]int
	}
	tests := []struct {
		name string
		args args
		want SubrectangleQueries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.rectangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubrectangleQueries_UpdateSubrectangle(t *testing.T) {
	type args struct {
		row1     int
		col1     int
		row2     int
		col2     int
		newValue int
	}
	tests := []struct {
		name string
		this *SubrectangleQueries
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SubrectangleQueries{}
			this.UpdateSubrectangle(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2, tt.args.newValue)
		})
	}
}

func TestSubrectangleQueries_GetValue(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		this *SubrectangleQueries
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SubrectangleQueries{}
			if got := this.GetValue(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("SubrectangleQueries.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
