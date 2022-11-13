package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Adder adds two numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
