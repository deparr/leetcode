func rob(nums []int) int {
	curr, prev := 0, 0

	// [curr, prev, v, v+1]
	for _, v := range nums {
		tmp := prev
		if v+curr > tmp {
			tmp = v + curr
		}

		curr = prev
		prev = tmp
	}

	return prev
}