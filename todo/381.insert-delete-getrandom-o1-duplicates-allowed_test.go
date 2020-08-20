// Generated TestRandomizedCollection_Insert
// Generated TestRandomizedCollection_Remove
// Generated TestRandomizedCollection_GetRandom
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want RandomizedCollection
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

func TestRandomizedCollection_Insert(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *RandomizedCollection
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedCollection{}
			if got := this.Insert(tt.args.val); got != tt.want {
				t.Errorf("RandomizedCollection.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedCollection_Remove(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *RandomizedCollection
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedCollection{}
			if got := this.Remove(tt.args.val); got != tt.want {
				t.Errorf("RandomizedCollection.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedCollection_GetRandom(t *testing.T) {
	tests := []struct {
		name string
		this *RandomizedCollection
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedCollection{}
			if got := this.GetRandom(); got != tt.want {
				t.Errorf("RandomizedCollection.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
