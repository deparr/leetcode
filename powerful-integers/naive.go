func powerfulIntegers(x int, y int, bound int) []int {
    currx, curry := 1, 1
    set := map[int]struct{} {}

    for currx <= bound {
        curry = 1
        for currx + curry <= bound {
            set[currx + curry] = struct{} {}
            curry *= y
            if curry == 1 { break }
        }
        currx *= x
        if currx == 1 { break }
    }

    res := make([]int, 0, len(set))
    for val, _ := range set {
        res = append(res, val)
	}

    return res
}
