func merge(nums1 []int, m int, nums2 []int, n int)  {
    mer := make([]int, m + n)

    aIdx, bIdx, mIdx := 0, 0, 0

    for aIdx < m || bIdx < n {
        nextVal := 0

        if bIdx >= n {
            nextVal = nums1[aIdx]
            aIdx++
        } else if aIdx >= m {
            nextVal = nums2[bIdx]
            bIdx++
        } else {
            if nums1[aIdx] < nums2[bIdx] {
               nextVal = nums1[aIdx] 
               aIdx++
            } else {
               nextVal = nums2[bIdx] 
               bIdx++
            }
        }

        mer[mIdx] = nextVal
        mIdx++
    }
    for i := range mer {
        nums1[i] = mer[i]
    }
}

