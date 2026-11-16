package g3ex1

import (
	"slices"
	"testing"
)

func TestG3ex1(t *testing.T) {
	got1, got2, got3 := g3ex1()
	want1 := []string{"Hello", "Hola"}
	want2 := []string{"Hola", "नमस्कार", "こんにちは"}
	want3 := []string{"こんにちは", "Привіт"}

	if (!slices.Equal(got1, want1)) || (!slices.Equal(got2, want2)) || (!slices.Equal(got3, want3)) {
		t.Errorf("got1 %q, got2 %q, got3 %q, wanted1 %q, wanted2 %q, wanted3 %q", got1, got2, got3, want1, want2, want3)
	}
}
