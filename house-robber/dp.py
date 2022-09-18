from typing import List

def rob(nums: List[int]) -> int:
    prev2, prev = 0, 0
    for val in nums:
        tmp = max(prev2 + val, prev)
        prev2 = prev
        prev = tmp
    return prev