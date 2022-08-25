package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)

	for i := len(cost) - 3; i > -1; i-- {
		min := cost[i+1]
		if cost[i+2] < min {
			min = cost[i+2]
		}
		cost[i] += min
	}

	if cost[1] < cost[0] {
		return cost[1]
	}

	return cost[0]

}

func main() {
	arr := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(arr))
}
