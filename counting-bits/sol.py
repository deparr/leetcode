class Solution:
    def countBits(self, n: int) -> List[int]:
        res = [0] * (n + 1)
        for curr in range(1, n + 1):
            if curr % 2 == 1:
                res[curr] = res[curr >> 1] + 1
            else:
                res[curr] = res[curr >> 1]
        return res
