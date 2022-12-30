func lengthOfLIS(nums []int) int {
    max := func(a []int) int {
        _max := a[0]
        for _, v := range a { if v > _max { _max = v }}
        return _max
    }

    increaseCount := make([]int, len(nums))

    increaseCount[0] = 1
    for i := 1; i < len(nums); i++ {
        maxes := make([]int, i)
        for j := i-1; j >= 0; j-- {
            if nums[i] > nums[j] {
                maxes[j] = increaseCount[j]
            }
        }
        increaseCount[i] = max(maxes) + 1
    }


    return max(increaseCount)
}
