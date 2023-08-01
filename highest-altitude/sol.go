func largestAltitude(gain []int) int {
    for i := 1; i < len(gain); i++ {
        gain[i] += gain[i-1]
    }

    max := gain[0]
    for i := 1; i < len(gain); i++ {
        if gain[i] > max {
            max = gain[i]
        }
    }

    if max < 0 {
        return 0
    }

    return max
}

