// Generated TestLFUCache_Get
// Generated TestLFUCache_Put
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LFUCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLFUCache_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *LFUCache
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LFUCache{}
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("LFUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLFUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		this *LFUCache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &LFUCache{}
			this.Put(tt.args.key, tt.args.value)
		})
	}
}
