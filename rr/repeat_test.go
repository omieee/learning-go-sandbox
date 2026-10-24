package iteration

import "testing"

func TestRepet(t *testing.T) {
	t.Run("when times is 3", func(t *testing.T) {
		got := Repeat("x", 3)
		want := "xxx"
		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
	t.Run("when times is 0", func(t *testing.T) {
		got := Repeat("b", 0)
		want := ""
		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})

	t.Run("when two chars are given and times is set to 3", func(t *testing.T) {
		got := Repeat("ab", 3)
		want := "ababab"
		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
}
