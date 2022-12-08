class Solution:
    def halveArray(self, nums: List[int]) -> int:
        #nums.sort(reverse=True)
        _sum = sum(nums)
        for i in range(len(nums)):
            nums[i] *= -1

        heapq.heapify(nums)

        reduction = 0
        ops = 0

        while _sum - reduction > _sum / 2:
            big = heapq.heappop(nums)
            big /= 2
            reduction += big * -1
            ops += 1
            heapq.heappush(nums, big)


        return ops
