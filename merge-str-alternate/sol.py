class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        res = []

        for a, b in zip(word1, word2):
            res.append(a)
            res.append(b)
        
        res = "".join(res)
        if (len(word1) - len(word2)) != 0:
            if len(word1) > len(word2):
                res += word1[len(word2):]
            else:
                res += word2[len(word1):]

        return res

