impl Solution {
    pub fn kth_factor(n: i32, k: i32) -> i32 {
        let mut count = 1;
        let mut factor = 1;

        while count < k {
            if factor == n { return -1 }
            factor += 1;
            while n % factor != 0 { factor += 1; }
            count += 1;
        }

        return factor;
    }
}
