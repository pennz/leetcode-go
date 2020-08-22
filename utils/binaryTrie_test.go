package utils

import (
	"reflect"
	"testing"
)

func TestInitBinaryTrie(t *testing.T) {
	tests := []struct {
		name string
		want *binaryTrie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitBinaryTrie(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitBinaryTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryTrie_insert(t *testing.T) {
	newBt := InitBinaryTrie()
	type fields struct {
		root *binaryTrieNode
	}
	type args struct {
		number int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"basic", fields{newBt.Root}, args{3}, 1},
		{"basic2", fields{newBt.Root}, args{3}, 1},
		{"basic3", fields{newBt.Root}, args{1 << 31}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &binaryTrie{
				root: tt.fields.root,
			}
			bt.Insert(tt.args.number)
			length := 0

			for _, v := range bt.root.childrens {
				if v != nil {
					length++
				}
			}

			if got := length; got != tt.want {
				t.Errorf("Insert function return = %v, want %v", got, tt.want)
			}
		})
	}
}
