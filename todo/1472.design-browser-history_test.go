Generated TestConstructor
Generated TestBrowserHistory_Visit
Generated TestBrowserHistory_Back
Generated TestBrowserHistory_Forward
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		homepage string
	}
	tests := []struct {
		name string
		args args
		want BrowserHistory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.homepage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Visit(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &BrowserHistory{}
			this.Visit(tt.args.url)
		})
	}
}

func TestBrowserHistory_Back(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &BrowserHistory{}
			if got := this.Back(tt.args.steps); got != tt.want {
				t.Errorf("BrowserHistory.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Forward(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &BrowserHistory{}
			if got := this.Forward(tt.args.steps); got != tt.want {
				t.Errorf("BrowserHistory.Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}
