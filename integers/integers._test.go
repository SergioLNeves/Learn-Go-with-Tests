package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("Should to sum A and B", func(t *testing.T) {
		sum := Add(1, 2)
		expected := 3

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
