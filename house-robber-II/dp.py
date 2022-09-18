from typing import List
def rob(self, nums: List[int]) -> int:
    def helper (nums: List[int]) -> int:
        prev2, prev = 0, 0
        for val in nums:
            tmp = max(prev2 + val, prev)
            prev2 = prev
            prev = tmp
        return prev
    return max(helper(nums[1:]), helper(nums[:-1]), nums[0])

print(rob(None, [10]))