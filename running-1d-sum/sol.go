func runningSum(nums []int) []int {
    sum := 0
    for i, val := range nums {
        sum += val
        nums[i] = sum
    }
    
    return nums
}
