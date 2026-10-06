package helloworld

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello!, World!"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Om Shanker")
	want := "Hello! Om Shanker"

	if got != want {
		t.Errorf("got %q and wanted %q", got, want)
	}
}

func TestHelloNameWithGreet(t *testing.T) {
	got := HelloNameWithGreet("Foxy")
	want := "Hello Foxy!"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
