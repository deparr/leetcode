func strStr(haystack string, needle string) (res int) {
	n := len(needle)
	ans = -1

	for i := 0; i < len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			ans = i
			break
		}
	}

	return
}
