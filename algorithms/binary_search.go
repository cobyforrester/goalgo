package algorithms

// Returns index of target if found, else -1
func BinarySearch(sortedArr []int, target int) int {
	return helper(sortedArr, target, 0, len(sortedArr) - 1)
}

// helper function for recursion
func helper(sortedArr []int, target int, leftIndex int, rightIndex int) int {
	if leftIndex == rightIndex && sortedArr[leftIndex] != target {
		return -1
	}
	index := (leftIndex + rightIndex) / 2
	indexValue := sortedArr[index]
	if indexValue == target {
		return index
	}

	if target < indexValue {
		return helper(sortedArr, target, leftIndex, index - 1)
	} else {
		return helper(sortedArr, target, index +1, rightIndex)
	}
}