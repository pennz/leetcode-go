// Generated TestProductOfNumbers_Add
// Generated TestProductOfNumbers_GetProduct
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want ProductOfNumbers
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

func TestProductOfNumbers_Add(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *ProductOfNumbers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{}
			this.Add(tt.args.num)
		})
	}
}

func TestProductOfNumbers_GetProduct(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		this *ProductOfNumbers
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ProductOfNumbers{}
			if got := this.GetProduct(tt.args.k); got != tt.want {
				t.Errorf("ProductOfNumbers.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
