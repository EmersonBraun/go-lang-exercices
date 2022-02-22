package array

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	verifySum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result '%d' expected '%d'", result, expected)
		}
	}
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("result '%d' expected '%d', data '%v'", result, expected, numbers)
		}
	})
	t.Run("Collection of any lenght", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("result '%d' expected '%d', data '%v'", result, expected, numbers)
		}
	})
	t.Run("Sum all", func(t *testing.T) {
		result := SumAll([]int{1,2}, []int{0,9})
		expected := []int{3,9}

		verifySum(t, result, expected)
	})
	t.Run("Sum rest", func(t *testing.T) {
		result := SumRest([]int{1,2}, []int{0,9})
		expected := []int{2,9}

		verifySum(t, result, expected)
	})
	t.Run("Sum some slices", func(t *testing.T) {
		result := SumRest([]int{1,2}, []int{0,9})
		expected := []int{2,9}

		verifySum(t, result, expected)
	})
	t.Run("Sum safe some slices", func(t *testing.T) {
		result := SumRest([]int{}, []int{3,4,5})
		expected := []int{0,9}

		verifySum(t, result, expected)
	})
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumRest(numbersToSum ...[]int) []int {
	var sum []int
    for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			final := numbers[1:]
			sum = append(sum, Sum(final))
		}
    }

    return sum
}