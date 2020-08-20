Generated TestConstructor
Generated TestMyCalendar_Book
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyCalendar
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

func TestMyCalendar_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		this *MyCalendar
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCalendar{}
			if got := this.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendar.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
