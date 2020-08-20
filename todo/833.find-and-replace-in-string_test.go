Generated Test_findReplaceString
package main

import "testing"

func Test_findReplaceString(t *testing.T) {
	type args struct {
		S       string
		indexes []int
		sources []string
		targets []string
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
			if got := findReplaceString(tt.args.S, tt.args.indexes, tt.args.sources, tt.args.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
