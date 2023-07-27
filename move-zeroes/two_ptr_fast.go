func moveZeroes(nums []int)  {
    l, r := 0, 0
    for r != len(nums) {
        if nums[r] != 0 {
            nums[l] = nums[r]
            l++
        }
        r++
    }

    for l != len(nums) {
        nums[l] = 0
        l++
    }
}

