func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool, len(candies))
    max := -1
    for _, v := range candies {
        if v > max {
            max = v
        }
    }

    for i, v := range candies {
        res[i] = (v + extraCandies) >= max
    }

    return res
}

