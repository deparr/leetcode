package main

import "fmt"

func rob(nums []int) int {
	prev_max, running_max := 0, 0

	// [prev_max, running_max, v, v+1]
	for _, v := range nums {
		tmp := running_max
		if v+prev_max > tmp {
			tmp = v + prev_max
		}

		prev_max = running_max
		running_max = tmp
	}

	return running_max
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
