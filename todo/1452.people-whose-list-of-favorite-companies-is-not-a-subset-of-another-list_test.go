Generated Test_peopleIndexes
package main

import (
	"reflect"
	"testing"
)

func Test_peopleIndexes(t *testing.T) {
	type args struct {
		favoriteCompanies [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peopleIndexes(tt.args.favoriteCompanies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peopleIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}
