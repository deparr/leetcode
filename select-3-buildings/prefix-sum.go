func numberOfWays(s string) int64 {
	n := len(s)
	count_one_zero := make([]int64, n)
	count_zero_one := make([]int64, n)

	var (zocount int64; count int64)
	for i := n-1; i >= 0; i-- {
		if s[i] == 0x30 {
			zocount++
		} else {
			count += zocount
		}
		count_one_zero[i] = count
	}


	zocount, count = 0, 0
	for i := n-1; i >= 0; i-- {
		if s[i] == 0x31 {
			zocount++
		} else {
			count += zocount
		}
		count_zero_one[i] = count
	}

	var result int64 = 0
	for i := 0; i < n; i++ {
		if s[i] == 0x30 {
			result += count_one_zero[i]
		} else {
			result += count_zero_one[i]
		}
	}

    return result
}
