// Generated TestMapSum_Insert
// Generated TestMapSum_Sum
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MapSum
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

func TestMapSum_Insert(t *testing.T) {
	type args struct {
		key string
		val int
	}
	tests := []struct {
		name string
		this *MapSum
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MapSum{}
			this.Insert(tt.args.key, tt.args.val)
		})
	}
}

func TestMapSum_Sum(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		this *MapSum
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MapSum{}
			if got := this.Sum(tt.args.prefix); got != tt.want {
				t.Errorf("MapSum.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
