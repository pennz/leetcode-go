package main

import "testing"

func Test_entityParser(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entityParser(tt.args.text); got != tt.want {
				t.Errorf("entityParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
