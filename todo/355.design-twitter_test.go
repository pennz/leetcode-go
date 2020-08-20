// Generated TestTwitter_PostTweet
// Generated TestTwitter_GetNewsFeed
// Generated TestTwitter_Follow
// Generated TestTwitter_Unfollow
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Twitter
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

func TestTwitter_PostTweet(t *testing.T) {
	type args struct {
		userId  int
		tweetId int
	}
	tests := []struct {
		name string
		this *Twitter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Twitter{}
			this.PostTweet(tt.args.userId, tt.args.tweetId)
		})
	}
}

func TestTwitter_GetNewsFeed(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name string
		this *Twitter
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Twitter{}
			if got := this.GetNewsFeed(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Twitter.GetNewsFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwitter_Follow(t *testing.T) {
	type args struct {
		followerId int
		followeeId int
	}
	tests := []struct {
		name string
		this *Twitter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Twitter{}
			this.Follow(tt.args.followerId, tt.args.followeeId)
		})
	}
}

func TestTwitter_Unfollow(t *testing.T) {
	type args struct {
		followerId int
		followeeId int
	}
	tests := []struct {
		name string
		this *Twitter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Twitter{}
			this.Unfollow(tt.args.followerId, tt.args.followeeId)
		})
	}
}
