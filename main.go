package main

import (
	"fmt"

	"github.com/cobyforrester/goalgo/algorithms"
)

func main() {
	sortedArr := []int{9, 9, 10}
	fmt.Println(algorithms.BinarySearch(sortedArr, 10))
}