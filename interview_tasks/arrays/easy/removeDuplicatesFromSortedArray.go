package easy

func removeDuplicates(array []int) ([]int, int) {

	pointer, minion := 0, 0

	for pointer < len(array)-1 {
		for array[pointer] == array[minion] {
			minion++
			if minion == len(array) {
				return array[:pointer+1], pointer + 1
			}
		}

		array[pointer+1] = array[minion]
		pointer++

	}

	return array[:pointer+1], pointer + 1
}
