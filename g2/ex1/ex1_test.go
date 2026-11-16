package ex1

import "testing"

func TestEx1(t *testing.T) {
	got1, got2 := ex1()
	want1 := 20
	want2 := 20.0

	if got1 != want1 && got2 != want2 {
		t.Errorf("got1 %d, got2 %g, wanted1 %d, wanted2 %g", got1, got2, want1, want2)
	}
}
