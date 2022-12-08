class Solution:
    def getAverages(self, nums: List[int], k: int) -> List[int]:
        # 0 - k inclusive
        # i -k i + k
        if k == 0:
            return nums

        if k > len(nums) // 2:
            return [-1 for _ in range(len(nums))]

        averages = [-1 for _ in range(k)]
        prefix = [nums[0]]
        for i in range(1, len(nums)):
            prefix.append(prefix[i-1] + nums[i])

        sumlen = 2 * k + 1
        if 2 * k < len(nums):
            averages.append(prefix[2*k] // sumlen)
        for i in range(k+1, len(nums)-k):
            averages.append((prefix[i+k] - prefix[i-k - 1 if i != k else 0]) // sumlen)

        for _ in range(k):
            averages.append(-1)

        return averages
