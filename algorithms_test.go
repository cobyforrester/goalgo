package main

import (
	"testing"

	"github.com/cobyforrester/goalgo/algorithms"
)

func TestBinarySearch(t *testing.T) {
	sortedArr := []int{1, 3, 5, 6, 6, 6, 7, 9, 9, 9, 10}

	actualResult := algorithms.BinarySearch(sortedArr, 5)
	expectedResult := 2

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}