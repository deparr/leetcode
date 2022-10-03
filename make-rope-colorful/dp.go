func minCost(colors string, neededTime []int) int {
    curr_min := 0
    last_index := -1
    for i := 0; i < len(colors); i++ {
        if i > 0 && colors[i] == colors[last_index] {
            if neededTime[i] < neededTime[last_index] {
                curr_min += neededTime[i]
                continue
            } else {
                curr_min += neededTime[last_index]
            }
        }
        last_index = i
    }
    return curr_min
}
