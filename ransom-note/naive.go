func canConstruct(ransomNote string, magazine string) bool {

	cache := map[rune]int{}

	for _, c := range magazine {
		cache[c] += 1
	}

	for _, c := range ransomNote {
		if cache[c] == 0 {
			return false
		}

		cache[c]--
	}

	return true

}