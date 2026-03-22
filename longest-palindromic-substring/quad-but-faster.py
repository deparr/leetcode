class Solution:
    def longestPalindrome(self, s: str) -> str:
        def find_palindrome_at(l, r):
            while l >= 0 and r < len(s) and s[l] == s[r]:
                l -= 1
                r += 1
            return l + 1, r - 1
        
        max_len = 0
        max_l, max_r = 0, 0

        for i in range(len(s) - 1):
            l, r = find_palindrome_at(i, i)
            new_len = r - l
            if new_len > max_len:
                max_len = new_len
                max_l = l
                max_r = r

            l, r = find_palindrome_at(i, i + 1)
            new_len = r - l
            if new_len > max_len:
                max_len = new_len
                max_l = l
                max_r = r
        
        return s[max_l:max_r + 1]
