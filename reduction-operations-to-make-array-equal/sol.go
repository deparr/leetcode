import "slices"

func reductionOperations(nums []int) int {
    slices.Sort(nums)
    ops := 0
    for lead := len(nums)-1; lead > 0; lead-- {
        if (nums[lead-1] != nums[lead]) {
            ops += len(nums) - lead
        }
    }

    return ops
}
