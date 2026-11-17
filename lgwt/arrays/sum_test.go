package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("when size is set to 5", func(t *testing.T) {
		// Previously when it was array we had got := Sum([5]int{1, 2, 3, 4, 5})
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15

		if got != want {
			t.Errorf("Got %d, Wanted %d", got, want)
		}
	})

	t.Run("when we have lesser size will not rely on array wil switch to slice", func(t *testing.T) {
		got := Sum([]int{1, 2, 3})
		want := 6

		if got != want {
			t.Errorf("Got %d, Wanted %d", got, want)
		}
	})

}
