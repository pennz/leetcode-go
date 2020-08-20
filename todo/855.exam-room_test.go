Generated TestConstructor
Generated TestExamRoom_Seat
Generated TestExamRoom_Leave
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want ExamRoom
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExamRoom_Seat(t *testing.T) {
	tests := []struct {
		name string
		this *ExamRoom
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ExamRoom{}
			if got := this.Seat(); got != tt.want {
				t.Errorf("ExamRoom.Seat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExamRoom_Leave(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		this *ExamRoom
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &ExamRoom{}
			this.Leave(tt.args.p)
		})
	}
}
