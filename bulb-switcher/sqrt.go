func bulbSwitch(n int) int {
    on := 0
    for i := 1; (i*i) <= n; i++ {
        on++
    }

    return on
}
