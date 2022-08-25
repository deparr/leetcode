func minCostClimbingStairs(cost []int) int {
	curr, next := 0, 0
	for i := len(cost) - 1; i > -1; i-- {
		min := curr
		if next < min {
			min = next
		}

		curr, next = cost[i]+min, curr
	}

	if next < curr {
		return next
	}

	return curr
}