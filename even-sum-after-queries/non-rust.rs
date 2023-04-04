impl Solution {
    pub fn sum_even_after_queries(mut nums: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut sum = 0;
        nums.iter().for_each(|x| if x % 2 == 0 {sum += x});
        let mut sums = vec![sum; queries.len()];

        for (i, query) in queries.iter().enumerate() {
            let qidx = query[1] as usize;
            let new_sum = nums[qidx] + query[0];
            let was_even = nums[qidx] % 2 == 0;

            if new_sum % 2 == 0 {
                if was_even {
                    sum += new_sum - nums[qidx];
                } else {
                    sum += new_sum
                }
            } else {
                if was_even {
                    sum -= nums[qidx];
                }
            }

            nums[qidx] = new_sum;

            sums[i] = sum;
        }

        return sums;
    }
}
