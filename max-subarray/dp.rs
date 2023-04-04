impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut curr_max = nums[0];
        let mut prev_sum = nums[0];
        for num in nums.iter().skip(1) {
            if prev_sum + *num > *num {
                prev_sum += *num;
            } else {
                prev_sum = *num;
            }

            if prev_sum > curr_max {
                curr_max = prev_sum;
            }
        }

       return curr_max;
    }
}
