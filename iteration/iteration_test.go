package iteration

import "testing"

func TestRepeat(t *testing.T) {
	checksRepeat := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected '%s', repeated '%s'", expected, repeated)
		}
	}

	t.Run("should it repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat(5, "a")
		expected := "aaaaa"
		checksRepeat(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(5, "a")
	}
}
