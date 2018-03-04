package main

import (
	"reflect"
	"testing"
)

func TestPages(t *testing.T) {
	tcs := []struct {
		n    int
		want []int
	}{
		{2, []int{1, 2}},
		{4, []int{1, 4, 2, 3}},
		{6, []int{1, 6, 2, 5, 3, 4}},
		{10, []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}},
	}
	for _, tc := range tcs {
		if got, want := pages(tc.n), tc.want; !reflect.DeepEqual(got, want) {
			t.Errorf("pages(%d): got %v, want %v", tc.n, got, want)
		}
	}
}
