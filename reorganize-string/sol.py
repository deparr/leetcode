class Solution:
    def reorganizeString(self, s: str) -> str:
        counts = [0] * 26
        for c in s:
            counts[ord(c) - ord('a')] += 1

        most_freq, most_idx = 0, 0

        for i, n in enumerate(counts):
            if n > most_freq:
                most_freq = n
                most_idx = i
        
        if most_freq > (len(s) + 1) / 2:
            return ""

        res = [''] * len(s)

        res_idx = 0
        for _ in range(most_freq):
            res[res_idx] = chr(ord('a') + most_idx)
            res_idx += 2
        counts[most_idx] = 0

        for letter_ord, c in enumerate(counts):
            for _ in range(c):
                if res_idx >= len(res):
                    res_idx = 1
                res[res_idx] = chr(ord('a') + letter_ord)
                res_idx += 2


        return "".join(res)
