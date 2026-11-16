package arrays

import (
	"slices"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5}, []int{}, []int{10})
	want := []int{2, 5, 14, 0, 0} // May be in the tutorial they missed the scenario when slice has only 1.. I decided to add 0

	if !slices.Equal(got, want) {
		t.Errorf("Got %v, Wanted, %v", got, want)
	}
}
