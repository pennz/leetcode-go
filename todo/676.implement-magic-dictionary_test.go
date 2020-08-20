// Generated TestMagicDictionary_BuildDict
// Generated TestMagicDictionary_Search
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MagicDictionary
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

func TestMagicDictionary_BuildDict(t *testing.T) {
	type args struct {
		dict []string
	}
	tests := []struct {
		name string
		this *MagicDictionary
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MagicDictionary{}
			this.BuildDict(tt.args.dict)
		})
	}
}

func TestMagicDictionary_Search(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		this *MagicDictionary
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MagicDictionary{}
			if got := this.Search(tt.args.word); got != tt.want {
				t.Errorf("MagicDictionary.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
