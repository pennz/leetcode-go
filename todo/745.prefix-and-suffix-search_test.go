Generated TestConstructor
Generated TestWordFilter_F
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want WordFilter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordFilter_F(t *testing.T) {
	type args struct {
		prefix string
		suffix string
	}
	tests := []struct {
		name string
		this *WordFilter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &WordFilter{}
			if got := this.F(tt.args.prefix, tt.args.suffix); got != tt.want {
				t.Errorf("WordFilter.F() = %v, want %v", got, tt.want)
			}
		})
	}
}
