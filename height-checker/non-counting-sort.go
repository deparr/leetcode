func heightChecker(heights []int) int {
    sorted := make([]int, len(heights))
    copy(sorted, heights)
    sort.Ints(sorted)
    fmt.Println(sorted)
    total := 0
    for i := range sorted {
        if sorted[i] != heights[i] {
            total++
        }
    }

    return total
    
}

