package utils

import (
	"reflect"
	"testing"
)

func TestInitTrie(t *testing.T) {
	tests := []struct {
		name string
		want *trie
	}{
		// TODO: Add test cases.
		{"basic", &trie{&trieNode{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitTrie(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitTrie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trie_insert(t *testing.T) {
	type fields struct {
		root *trieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"basic", fields{InitTrie().root}, args{"a"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trieInstance := &trie{
				root: tt.fields.root,
			}
			trieInstance.insert(tt.args.word)
			length := 0
			for _, v := range trieInstance.root.childrens {
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

func Test_trie_find(t *testing.T) {
	type fields struct {
		root *trieNode
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t := &trie{
			//	root: tt.fields.root,
			//}
			//if got := t.find(tt.args.word); got != tt.want {
			//	t.Errorf("trie.find() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
