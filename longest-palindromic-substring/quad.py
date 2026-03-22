class Solution:
    def longestPalindrome(self, s: str) -> str:
        max_sub = s[0]
        for l in range(len(s)):
            if len(s) - l <= len(max_sub):
                break
            r = len(s) - 1
            while l < r:
                while s[r] != s[l]:
                    r -= 1
                sub = s[l:r+1]
                if len(sub) > len(max_sub) and sub == sub[::-1]:
                    max_sub = sub
                r -= 1

        return max_sub
