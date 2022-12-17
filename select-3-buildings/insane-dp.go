func  numberOfWays(s string) int64 {
	var nways int64 = 0
	var one_pairs int64 = 0
	var zero_pairs int64 = 0
	var ones int64 = 0
	var zeros int64 = 0

	for i:=0; i<len(s); i++ {
		if s[i] == 0x31 {
			zeros++
			zero_pairs += ones
			nways += one_pairs
		} else {
			ones++
			one_pairs += zeros
			nways += zero_pairs
		}
	}

	return nways
}
