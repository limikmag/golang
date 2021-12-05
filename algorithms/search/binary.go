package search

// Binary search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// This function uses recursive call to itself.
// If a target is found, the index of the target is returned. Else the function return -1
func Binary(sortedSlice []int, target int, start int, end int) int {
	if start > end || len(sortedSlice) == 0 {
		return -1
	}

	mid := int(start + (end-start)/2)

	if target > sortedSlice[mid] {
		return Binary(sortedSlice, target, start, mid-1)
	} else if target < sortedSlice[mid] {
		return Binary(sortedSlice, target, mid+1, end)
	} else {
		return mid
	}
}

// BinaryIterative search for target within a sorted array by repeatedly dividing the array in half and comparing the midpoint with the target.
// Unlike Binary, this function uses iterative method and not recursive.
// If a target is found, the index of the target is returned. Else the function return -1
func BinaryIterative(sortedSlice []int, target int, start int, end int) int {
	startIndex := start
	endIndex := end
	var mid int

	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if sortedSlice[mid] > target {
			endIndex = mid - 1
		} else if sortedSlice[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
