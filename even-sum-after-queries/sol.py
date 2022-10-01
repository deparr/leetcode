class Solution:
    def sumEvenAfterQueries(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        ans = list()
        even_sum = sum(val for val in nums if val % 2 == 0)

        for val, idx in queries:
            if nums[idx] % 2 == 0:
                even_sum -= nums[idx]
            
            nums[idx] += val
            
            if nums[idx] % 2 == 0:
                even_sum += nums[idx]
            
            ans.append(even_sum)
            
        return ans
