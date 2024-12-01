package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, n := range nums {
		if idx, exists := numMap[target-n]; exists {
			return []int{i, idx}
		}

		numMap[n] = i
	}

	return []int{}
}
