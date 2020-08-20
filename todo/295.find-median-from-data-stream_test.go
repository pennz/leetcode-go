Generated TestConstructor
Generated TestMedianFinder_AddNum
Generated TestMedianFinder_FindMedian
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MedianFinder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedianFinder_AddNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *MedianFinder
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MedianFinder{}
			this.AddNum(tt.args.num)
		})
	}
}

func TestMedianFinder_FindMedian(t *testing.T) {
	tests := []struct {
		name string
		this *MedianFinder
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MedianFinder{}
			if got := this.FindMedian(); got != tt.want {
				t.Errorf("MedianFinder.FindMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
