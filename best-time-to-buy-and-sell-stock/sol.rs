impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut profit = 0;
        let mut buy_price = prices[0];
        for i in 1..prices.len() {
            if (prices[i] < buy_price) {
                buy_price = prices[i];
            } else {
                profit = std::cmp::max(prices[i] - buy_price, profit)
            }
        }
        
        return profit;
    }
}
