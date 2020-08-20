Generated TestConstructor
Generated TestMyCalendarTwo_Book
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyCalendarTwo
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

func TestMyCalendarTwo_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		this *MyCalendarTwo
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCalendarTwo{}
			if got := this.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendarTwo.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
