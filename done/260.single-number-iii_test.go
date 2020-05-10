package main

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{12, 21}}, []int{12, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberMap(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{12, 21}}, []int{12, 21}},
		{"", args{[]int{12, 1, 1, 21}}, []int{21, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want_alt := []int{tt.want[1], tt.want[0]}
			if got := singleNumberMap(tt.args.nums); !(reflect.DeepEqual(got, tt.want) || reflect.DeepEqual(got, want_alt)) {
				t.Errorf("singleNumberMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
