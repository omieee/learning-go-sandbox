package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatSuccess(t *testing.T) {
	t.Run("when given character is 'x' and times is 3", func(t *testing.T) {
		got := Repeat("x", 3)
		want := "xxx"

		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})

	t.Run("when character is 'b' and times is 0", func(t *testing.T) {
		got := Repeat("b", 0)
		want := ""

		if want != got {
			t.Errorf("wanted %q got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("q", 3)
	}
}

func ExampleRepeat() {
	got := Repeat("w", 3)
	fmt.Println(got)

	//Output: www
}
