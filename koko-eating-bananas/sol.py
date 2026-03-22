class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        hi = max(piles)
        if len(piles) == h:
            return hi
        lo = 1
        min_eat = hi
        while lo < hi:
            mid = lo  + (hi - lo) // 2
            hours_taken = 0
            for n in piles:
                hours_taken += (n + mid - 1) // mid
            if hours_taken <= h:
                min_eat = min(mid, min_eat)
                hi = mid
            else:
                lo = mid + 1

        
        return min_eat
