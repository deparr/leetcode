func findDifference(nums1 []int, nums2 []int) [][]int {
    set := map[int]struct{} {}
    set2 := map[int]struct{} {}
    res := make([][]int, 2)
    for _, val := range nums1 {
        set[val] = struct {} {}
    }

    for _, val := range nums2 {
        set2[val] = struct {} {}
    }

    for val, _ := range set2 { 
        _, prs := set[val]
        if !prs {
            res[1] = append(res[1], val)
        }

    }

    for val, _ := range set { 
        _, prs := set2[val]
        if !prs {
            res[0] = append(res[0], val)
        }

    }

    return res
}
