package integer

import "testing"

func TestAdd(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d', result '%d'", expected, sum)
    }
}

func Add(x, y int) int {
	return x + y
}
