Generated Test_watchedVideosByFriends
package main

import (
	"reflect"
	"testing"
)

func Test_watchedVideosByFriends(t *testing.T) {
	type args struct {
		watchedVideos [][]string
		friends       [][]int
		id            int
		level         int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := watchedVideosByFriends(tt.args.watchedVideos, tt.args.friends, tt.args.id, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("watchedVideosByFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
