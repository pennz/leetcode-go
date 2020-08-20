Generated TestConstructor
Generated TestWordDictionary_AddWord
Generated TestWordDictionary_Search
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want WordDictionary
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

func TestWordDictionary_AddWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *WordDictionary
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &WordDictionary{}
			this.AddWord(tt.args.word)
		})
	}
}

func TestWordDictionary_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *WordDictionary
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &WordDictionary{}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("WordDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
