func findMaxAverage(nums []int, k int) float64 {
    total := 0
    for i := 0; i < k; i++ {
        total += nums[i]
    }
    max := total


    left := nums[0]
    for i := k; i < len(nums); i++ {
        total += nums[i]
        total -= left

        if total > max {
            max = total
        }
        left = nums[i-k+1]
    }

    return float64(max) / float64(k)
}

