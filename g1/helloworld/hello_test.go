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
	t.Run("when name is available", func(t *testing.T) {
		got := HelloNameWithGreet("Foxy", "french")
		want := "Bonjour! Foxy"
		assertCorrectMessage(t, got, want)
	})
	t.Run("when name is not available, return World!", func(t *testing.T) {
		got := HelloNameWithGreet("", "spanish")
		want := "Hola! World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := HelloNameWithGreet("Om Shanker", "spanish")
		want := "Hola! Om Shanker"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
