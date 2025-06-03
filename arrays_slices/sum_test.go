package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	// t.Run("test for variable array/slices", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	// 	}
	// })
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make the sum of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 3, 5})
		want := []int{2, 8}

		checkSums(t, got, want)
	})

	t.Run("safely make the sum of tails for empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 3, 5})
		want := []int{0, 8}

		checkSums(t, got, want)
	})
}
