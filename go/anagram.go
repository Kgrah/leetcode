package main

func isAnagram(s, t string) bool {
	abc := 26

	charSetS := make([]int, abc)
	charSetT := make([]int, abc)

	for _, r := range s {
		charSetS[charPosition(r)]++
	}

	for _, r := range t {
		charSetT[charPosition(r)]++
	}

	if len(charSetS) != len(charSetT) {
		return false
	}

	for i, c := range charSetS {
		if charSetT[i] != c {
			return false
		}
	}

	return true
}

func charPosition(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a')
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A')
	}

	return -1
}
