package iteration

import "testing"

const numberOfRepetitions = 5
func TestIteration(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected '%s' result '%s'", expected, result)
	}
}

func Repeat(char string) string {
	var rep string
    for i := 0; i < numberOfRepetitions; i++ {
        rep = rep + char
    }
    return rep
}