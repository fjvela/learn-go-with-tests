package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		checkTestResult(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := Sum(numbers)
		want := 45

		checkTestResult(t, got, want)
	})

}

func TestSumAll(t *testing.T) {

	t.Run("sum of two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkSliceTestResult(t, got, want)
	})

	t.Run("sum of any size slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9, 10})
		want := []int{6, 19}

		checkSliceTestResult(t, got, want)
	})

}

func TestSumAllTains(t *testing.T) {

	checkSums := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if len(got) != len(want) {
			t.Errorf("got %d slices, want %d slices", len(got), len(want))
		}
	}
	t.Run("sum of any size slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 10})
		want := []int{5, 19}

		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 10})
		want := []int{0, 19}

		checkSums(t, got, want)
	})
}

func checkSliceTestResult(t *testing.T, got []int, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkTestResult(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
