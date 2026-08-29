class Solution:
    def countGoodStrings(self, low: int, high: int, zero: int, one: int) -> int:
        mod_lim: int = 10 ** 9 + 7 
        cache = [0] * (high + 1)
        cache[0] = 1

        for i in range(1, high + 1):
            if i >= zero:
                cache[i] += cache[i - zero]
            if i >= one:
                cache[i] += cache[i - one]
            cache[i] %= mod_lim

        return sum(cache[low:]) % mod_lim
