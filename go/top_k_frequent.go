package main

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, n := range nums {
		freqMap[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for n, freq := range freqMap {
		buckets[freq] = append(buckets[freq], n)
	}

	var result []int
	for i := len(buckets) - 1; i > 0 && len(result) < k; i-- {
		for _, n := range buckets[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
