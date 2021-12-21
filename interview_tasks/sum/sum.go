package sum

// find in O(n) time complexity if there are 2 numbers in slice which will sum to X
func Sum(slice []int, sumNum int) bool {
	helper := make(map[int]int)

	for _, num := range slice {
		diff := sumNum - num
		if _, ok := helper[diff]; ok {
			return true
		} else {
			helper[num] = diff
		}
	}

	return false
}
