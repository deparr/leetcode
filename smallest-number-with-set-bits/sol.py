class Solution:
    def smallestNumber(self, n: int) -> int:
        factor = int(math.log2(n)) + 1
        return (1 << factor) - 1
