class Solution:
    def getDescentPeriods(self, prices: List[int]) -> int:
        total_periods = 1
        period_start = 0

        for i in range(1, len(prices)):
            total_periods += 1
            if prices[i-1] - prices[i] != 1:
                period_start = i
            else:
                total_periods += i - period_start

        return total_periods
