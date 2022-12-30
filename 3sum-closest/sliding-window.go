func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    minDiffSum := nums[0] + nums[1] + nums[len(nums) - 1]
    minDiff := minDiffSum - target
    if minDiff < 0 {
        minDiff *= -1
    }

    for k := 0; k < len(nums)-2; k++ {
        i := k + 1
        j := len(nums) - 1

        for i < j {
            sum := nums[i] + nums[j] + nums[k]

            if sum == target {
                return sum
            }

            if sum < target {
                i++
            } else {
                j--
            }

            currdiff := sum - target
            if currdiff < 0 {
                currdiff *= -1
            }
            if currdiff < minDiff {
                minDiff = currdiff
                minDiffSum = sum
            }

        }
    }

    return minDiffSum
    
}
