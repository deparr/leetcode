class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        res = []
        _sum = 0
        
        for num in nums:
            _sum += num
            res.append(_sum)
        return res
