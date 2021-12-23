package easy

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]


TimeComplexity - O(n)
SpaceComplexity - O(n)

*/

func twoSum(list []int, target int) []int {
	helper := make(map[int]int)

	for i := 0; i < len(list); i++ {
		diff := target - list[i]

		if _, ok := helper[diff]; ok {
			return []int{helper[diff], i}

		}
		helper[list[i]] = i

	}

	return nil
}
