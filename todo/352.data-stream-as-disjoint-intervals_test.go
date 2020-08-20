Generated TestConstructor
Generated TestSummaryRanges_AddNum
Generated TestSummaryRanges_GetIntervals
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want SummaryRanges
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

func TestSummaryRanges_AddNum(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *SummaryRanges
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SummaryRanges{}
			this.AddNum(tt.args.val)
		})
	}
}

func TestSummaryRanges_GetIntervals(t *testing.T) {
	tests := []struct {
		name string
		this *SummaryRanges
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &SummaryRanges{}
			if got := this.GetIntervals(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryRanges.GetIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
