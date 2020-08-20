Generated TestConstructor
Generated TestCashier_GetBill
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n        int
		discount int
		products []int
		prices   []int
	}
	tests := []struct {
		name string
		args args
		want Cashier
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n, tt.args.discount, tt.args.products, tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCashier_GetBill(t *testing.T) {
	type args struct {
		product []int
		amount  []int
	}
	tests := []struct {
		name string
		this *Cashier
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Cashier{}
			if got := this.GetBill(tt.args.product, tt.args.amount); got != tt.want {
				t.Errorf("Cashier.GetBill() = %v, want %v", got, tt.want)
			}
		})
	}
}
