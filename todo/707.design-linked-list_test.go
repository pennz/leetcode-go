Generated TestConstructor
Generated TestMyLinkedList_Get
Generated TestMyLinkedList_AddAtHead
Generated TestMyLinkedList_AddAtTail
Generated TestMyLinkedList_AddAtIndex
Generated TestMyLinkedList_DeleteAtIndex
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyLinkedList
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

func TestMyLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{}
			if got := this.Get(tt.args.index); got != tt.want {
				t.Errorf("MyLinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{}
			this.AddAtHead(tt.args.val)
		})
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{}
			this.AddAtTail(tt.args.val)
		})
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{}
			this.AddAtIndex(tt.args.index, tt.args.val)
		})
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyLinkedList{}
			this.DeleteAtIndex(tt.args.index)
		})
	}
}
