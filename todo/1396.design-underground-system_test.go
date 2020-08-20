Generated TestConstructor
Generated TestUndergroundSystem_CheckIn
Generated TestUndergroundSystem_CheckOut
Generated TestUndergroundSystem_GetAverageTime
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want UndergroundSystem
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

func TestUndergroundSystem_CheckIn(t *testing.T) {
	type args struct {
		id          int
		stationName string
		t           int
	}
	tests := []struct {
		name string
		this *UndergroundSystem
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &UndergroundSystem{}
			this.CheckIn(tt.args.id, tt.args.stationName, tt.args.t)
		})
	}
}

func TestUndergroundSystem_CheckOut(t *testing.T) {
	type args struct {
		id          int
		stationName string
		t           int
	}
	tests := []struct {
		name string
		this *UndergroundSystem
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &UndergroundSystem{}
			this.CheckOut(tt.args.id, tt.args.stationName, tt.args.t)
		})
	}
}

func TestUndergroundSystem_GetAverageTime(t *testing.T) {
	type args struct {
		startStation string
		endStation   string
	}
	tests := []struct {
		name string
		this *UndergroundSystem
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &UndergroundSystem{}
			if got := this.GetAverageTime(tt.args.startStation, tt.args.endStation); got != tt.want {
				t.Errorf("UndergroundSystem.GetAverageTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
