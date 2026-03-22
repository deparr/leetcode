class Solution:
    def jump(self, nums: List[int]) -> int:
        cache = [10001] * len(nums)
        cache[0] = 0
        for i, n in enumerate(nums):
            jump_range = min(len(nums), i + n + 1)
            for j in range(i + 1, jump_range):
                cache[j] = min(cache[j], cache[i] + 1)
        return cache[-1]
        
