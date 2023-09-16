func findMin(nums []int) int {
    left, right := 0, len(nums) - 1
    mid := (left + right) / 2
    res := nums[0]
    for left <= right {
        if nums[mid] >= nums[0] {
            left = mid + 1
        } else {
            res = nums[mid]
            right = mid - 1
        }
        mid = (left + right) / 2
    }
    return res
}

