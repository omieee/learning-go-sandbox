package ex2

import "testing"

func TestEx1(t *testing.T) {
	got1, got2 := ex2()
	want1 := 10
	want2 := 10.0

	if got1 != want1 && got2 != want2 {
		t.Errorf("got1 %d, got2 %g, wanted1 %d, wanted2 %g", got1, got2, want1, want2)
	}
}
