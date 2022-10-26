func maxSubArray(nums []int) int {
    curr_max := nums[0]
    prev_sum := nums[0]

    for i := 1; i < len(nums); i++ {
        if prev_sum + nums[i] > nums[i] {
            prev_sum += nums[i]
        } else {
            prev_sum = nums[i]
        }

        if prev_sum > curr_max {
            curr_max = prev_sum
        }
    }

    return curr_max
}
