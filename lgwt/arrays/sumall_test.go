package arrays

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	want := []int{3, 6, 15}

	if !slices.Equal(got, want) {
		t.Errorf("Got %v, Wanted %v", got, want)
	}
}
