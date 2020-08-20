package main

import "testing"

func Test_findSecretWord(t *testing.T) {
	type args struct {
		wordlist []string
		master   *Master
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findSecretWord(tt.args.wordlist, tt.args.master)
		})
	}
}
