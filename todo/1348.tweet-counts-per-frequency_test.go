// Generated TestTweetCounts_RecordTweet
// Generated TestTweetCounts_GetTweetCountsPerFrequency
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want TweetCounts
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

func TestTweetCounts_RecordTweet(t *testing.T) {
	type args struct {
		tweetName string
		time      int
	}
	tests := []struct {
		name string
		this *TweetCounts
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TweetCounts{}
			this.RecordTweet(tt.args.tweetName, tt.args.time)
		})
	}
}

func TestTweetCounts_GetTweetCountsPerFrequency(t *testing.T) {
	type args struct {
		freq      string
		tweetName string
		startTime int
		endTime   int
	}
	tests := []struct {
		name string
		this *TweetCounts
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TweetCounts{}
			if got := this.GetTweetCountsPerFrequency(tt.args.freq, tt.args.tweetName, tt.args.startTime, tt.args.endTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TweetCounts.GetTweetCountsPerFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
