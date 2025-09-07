func sumZero(n int) []int {
    res := make([]int, n)
    for i, n2 := 0, n / 2; i < n - 1; {
        res[i] = n2 * -1
        res[i + 1] = n2
        n2--
        i += 2
    }

    if n % 2 == 1 {
        res[n-1] = 0
    }
    
    return res
}

