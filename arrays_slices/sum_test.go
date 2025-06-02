package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	got := SumAll([]int{1,2}, []int{1,3,5})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// t.Run("test for variable array/slices", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := Sum(numbers)
	// 	want := 6

	// 	if got != want {
	// 		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	// 	}
	// })
}
