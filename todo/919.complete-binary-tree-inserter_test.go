// Generated TestCBTInserter_Insert
// Generated TestCBTInserter_Get_root
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want CBTInserter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCBTInserter_Insert(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		this *CBTInserter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CBTInserter{}
			if got := this.Insert(tt.args.v); got != tt.want {
				t.Errorf("CBTInserter.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCBTInserter_Get_root(t *testing.T) {
	tests := []struct {
		name string
		this *CBTInserter
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CBTInserter{}
			if got := this.Get_root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CBTInserter.Get_root() = %v, want %v", got, tt.want)
			}
		})
	}
}
