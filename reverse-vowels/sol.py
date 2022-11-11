class Solution:
    def reverseVowels(self, s: str) -> str:
        frontVowel, backVowel = 0, len(s)-1
        vowels = "aeiouAEIOU"
        slist = list(s)
        while frontVowel < backVowel:
            while slist[frontVowel] not in vowels and frontVowel < backVowel:
                frontVowel += 1
            while slist[backVowel] not in vowels and backVowel > frontVowel:
                backVowel -= 1
            slist[frontVowel], slist[backVowel] = slist[backVowel], slist[frontVowel]
            frontVowel += 1
            backVowel -= 1

        return ''.join(slist)
