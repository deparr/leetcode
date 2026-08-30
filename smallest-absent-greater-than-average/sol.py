class Solution:
    def smallestAbsent(self, nums: List[int]) -> int:
        avg = sum(nums) // len(nums)
        num_set = set(nums)
        diff = set([x - avg for x in nums])
        if avg < 0:
            for i in range(1, 101):
                if i not in num_set:
                    return i
        else:
            for i in range(1, 101):
                if i not in diff:
                    return i + avg

