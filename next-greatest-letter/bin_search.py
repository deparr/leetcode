class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        target_num = ord(target)
        mid = None
        lo, hi = 0, len(letters) - 1
        if target_num >= ord(letters[hi]):
            return letters[0]

        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if ord(letters[mid]) <= target_num:
                lo = mid + 1
            else:
                hi = mid - 1

        return letters[lo]
