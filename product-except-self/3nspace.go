func productExceptSelf(nums []int) []int {
    pre, post := make([]int, len(nums)), make([]int, len(nums))
    prod := make([]int, len(nums))
    
    pre[0], post[len(nums)-1] = nums[0], nums[len(nums)-1] 
    
    for i := 1; i < len(nums); i++ {
        pre[i] = nums[i] * pre[i-1]
        post[len(post)-i-1] = nums[len(nums)-i-1] * post[len(post)-i]
    }
        
    for i := range prod {
        if i == 0 {
            prod[0] = post[1]
        } else if i == len(prod) -1 {
            prod[i] = pre[i-1]
        } else {
            prod[i] = post[i+1] * pre[i-1]
        }
    }

    return prod
    
}
