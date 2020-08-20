// Generated TestMyCircularQueue_EnQueue
// Generated TestMyCircularQueue_DeQueue
// Generated TestMyCircularQueue_Front
// Generated TestMyCircularQueue_Rear
// Generated TestMyCircularQueue_IsEmpty
// Generated TestMyCircularQueue_IsFull
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
		want MyCircularQueue
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

func TestMyCircularQueue_EnQueue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *MyCircularQueue
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.EnQueue(tt.args.value); got != tt.want {
				t.Errorf("MyCircularQueue.EnQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.DeQueue(); got != tt.want {
				t.Errorf("MyCircularQueue.DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Front(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.Front(); got != tt.want {
				t.Errorf("MyCircularQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Rear(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.Rear(); got != tt.want {
				t.Errorf("MyCircularQueue.Rear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsFull(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularQueue{}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("MyCircularQueue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
