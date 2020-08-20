Generated TestConstructor
Generated TestTimeMap_Set
Generated TestTimeMap_Get
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want TimeMap
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

func TestTimeMap_Set(t *testing.T) {
	type args struct {
		key       string
		value     string
		timestamp int
	}
	tests := []struct {
		name string
		this *TimeMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TimeMap{}
			this.Set(tt.args.key, tt.args.value, tt.args.timestamp)
		})
	}
}

func TestTimeMap_Get(t *testing.T) {
	type args struct {
		key       string
		timestamp int
	}
	tests := []struct {
		name string
		this *TimeMap
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TimeMap{}
			if got := this.Get(tt.args.key, tt.args.timestamp); got != tt.want {
				t.Errorf("TimeMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
