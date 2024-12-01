package main

// bucket approach of creating a map of frequency buckets
// then calculating the char array map for each string
// then using the char array map as the map key and appending
// the string to the bucket correlating to that char array map

func groupAnagrams(strs []string) [][]string {
	freqs := make(map[[26]int][]string)

	for _, s := range strs {
		var f [26]int

		for _, r := range s {
			f[charPosition(r)-'a']++
		}

		if strS, exists := freqs[f]; exists {
			strS = append(strS, s)
			freqs[f] = strS
		} else {
			freqs[f] = []string{s}
		}
	}

	result := make([][]string, 0)
	for _, strS := range freqs {
		r := make([]string, 0)
		r = append(r, strS...)
		result = append(result, r)
	}

	return result
}
