func removeDuplicates(nums []int) int {
    borders := make([]int, 1, 30)

    prev := nums[0]
    borders[0] = 0
    for i := 1; i < len(nums); i++ {
        if prev != nums[i] {
            borders = append(borders, i)
        }
        prev = nums[i]
    }

    fmt.Println(borders)
    for i, b := range borders {
        nums[i] = nums[b]
    }

    return len(borders)
}
