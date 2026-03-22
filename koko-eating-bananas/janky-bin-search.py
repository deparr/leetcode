class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        hi = max(piles)
        if len(piles) == h:
            return hi
        lo = 0
        min_eat = hi
        while lo < hi:
            mid = (hi + lo + 1) // 2
            if mid == min_eat:
                break
            hours_taken = sum([n // mid + int(n % mid > 0) for n in piles])
            if hours_taken <= h:
                min_eat = min(mid, min_eat)
                hi = mid
            else:
                lo = mid

        
        return min_eat


