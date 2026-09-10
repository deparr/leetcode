class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        target_num = ord(target)
        for x in letters:
            if ord(x) > target_num:
                return x
        return letters[0]
