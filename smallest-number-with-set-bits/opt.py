class Solution:
    def smallestNumber(self, n: int) -> int:
        return (1 << n.bit_lenth()) - 1
