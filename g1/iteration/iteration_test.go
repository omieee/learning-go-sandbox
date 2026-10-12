package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"

	if want != got {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}

func ExampleRepeat() {
	got := Repeat("a")
	fmt.Println(got)

	//Output: aaaaa
}
