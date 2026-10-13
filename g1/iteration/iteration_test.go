package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatSuccess(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"

	if want != got {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func TestRepeatFail(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaaa"

	if want == got {
		t.Errorf("wanted %q got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	got := Repeat("a", 3)
	fmt.Println(got)

	//Output: aaa
}
