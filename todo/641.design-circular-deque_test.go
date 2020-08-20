// Generated TestMyCircularDeque_InsertFront
// Generated TestMyCircularDeque_InsertLast
// Generated TestMyCircularDeque_DeleteFront
// Generated TestMyCircularDeque_DeleteLast
// Generated TestMyCircularDeque_GetFront
// Generated TestMyCircularDeque_GetRear
// Generated TestMyCircularDeque_IsEmpty
// Generated TestMyCircularDeque_IsFull
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want MyCircularDeque
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_InsertFront(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *MyCircularDeque
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.InsertFront(tt.args.value); got != tt.want {
				t.Errorf("MyCircularDeque.InsertFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_InsertLast(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *MyCircularDeque
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.InsertLast(tt.args.value); got != tt.want {
				t.Errorf("MyCircularDeque.InsertLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_DeleteFront(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.DeleteFront(); got != tt.want {
				t.Errorf("MyCircularDeque.DeleteFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_DeleteLast(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.DeleteLast(); got != tt.want {
				t.Errorf("MyCircularDeque.DeleteLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_GetFront(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.GetFront(); got != tt.want {
				t.Errorf("MyCircularDeque.GetFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_GetRear(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.GetRear(); got != tt.want {
				t.Errorf("MyCircularDeque.GetRear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularDeque.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_IsFull(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularDeque
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("MyCircularDeque.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
