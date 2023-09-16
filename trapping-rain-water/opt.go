func trap(height []int) int {
    trappedWater := 0
    left, right := 0, len(height) - 1
    leftMax, rightMax := height[left], height[right]

    for left < right {
        if leftMax < rightMax {
            left++
            if height[left] > leftMax {
                leftMax = height[left]
            }
            trappedWater += leftMax - height[left]
        } else {
            right--
            if height[right] > rightMax {
                rightMax = height[right]
            }
            trappedWater += rightMax - height[right]
        }
    }

    return trappedWater
}
