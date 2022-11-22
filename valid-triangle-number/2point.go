func triangleNumber(nums []int) int {
    sort.Slice(nums, func(x, y int) bool { return nums[x] < nums[y] } )
    count := 0

    for k := 2; k < len(nums); k++ {
        i := 0
        j := k - 1
        for i < j {
            if nums[i] +  nums[j] > nums[k] {
                count += j - i
                j--
            } else {
                i++
            }
        }
    }

    return count
    
}
