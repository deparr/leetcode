func moveZeroes(nums []int)  {
    end := len(nums)-1
    for i := len(nums)-1; i >= 0; i-- {
       if nums[i] == 0 {
           nums[end], nums[i] = nums[i], nums[end]
           end--
           j := i
           for j < end {
               nums[j], nums[j+1] = nums[j+1], nums[j]
               j++
           }
       }
    }
}

