func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxArea(height []int) int {

    res := 0
    left, right := 0, len(height)-1
    for left < right {
        if height[right] > height[left] {
            res = max(res, (right-left) * height[left])
            left++
        } else {
            res = max(res, (right-left) * height[right])
            right--
        }
    }

    return res
}
