impl Solution {
    pub fn min_operations_to_make_median_k(mut nums: Vec<i32>, k: i32) -> i64 {
        nums.sort_unstable();

        let mid = nums.len() / 2;
        let (median, median_idx) = if nums.len() % 2 == 0 {
            if nums[mid] >= nums[mid - 1] { (nums[mid], mid) }
            else { (nums[mid - 1], mid - 1) }
        } else {
            (nums[nums.len() / 2], nums.len() / 2)
        };

        if median == k {
            return 0;
        }

        let mut res: i64 = 0;

        let mut i = median_idx;
        if k < median {
            while nums[i] > k {
                res += (nums[i] - k) as i64;
                if i == 0 { break }
                i -= 1;
            }
        } else {
            while i < nums.len() && nums[i] < k {
                res += (k - nums[i]) as i64;
                i += 1;
            }
        }

        return res;
    }
}
