package main

import "testing"

func Test_longestDiverseString(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		//{"basic", args{1, 1, 7}, "ccbccacc"},
		//{"basic2", args{2, 2, 1}, "aabbc"},
		{"basic3", args{7, 1, 0}, "aabaa"},
		{"corner1", args{1, 0, 0}, "a"},
		/* Example 1:
		 *
		 *
		 * Input: a = 1, b = 1, c = 7
		 * Output: "ccaccbcc"
		 * Explanation: "ccbccacc" would also be a correct answer.
		 *
		 *
		 * Example 2:
		 *
		 *
		 * Input: a = 2, b = 2, c = 1
		 * Output: "aabbc"
		 *
		 *
		 * Example 3:
		 *
		 *
		 * Input: a = 7, b = 1, c = 0
		 * Output: "aabaa"
		 * Explanation: It's the only correct answer in this case.
		 */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDiverseString(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
