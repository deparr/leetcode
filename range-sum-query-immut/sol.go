type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    summed := make([]int, len(nums))
    sum := nums[0]
    summed[0] = sum
    for i := 1; i < len(nums); i++ {
        sum += nums[i]
        summed[i] = sum
    }
    return NumArray{summed}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left > 0 {
        return this.nums[right] - this.nums[left - 1]
    }
    return this.nums[right]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
