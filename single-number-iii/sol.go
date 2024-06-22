func singleNumber(nums []int) []int {
    if len(nums) == 2 {
        return nums
    }

	cache := 0
    for _, n := range nums {
        cache ^= n
    } 

    lowerBit := cache & -cache
    result := make([]int, 2)
    for _, n := range nums {
        if lowerBit & n == 0 {
            result[0] ^= n
        } else {
            result [1] ^= n
        }
    }

    return result
}
