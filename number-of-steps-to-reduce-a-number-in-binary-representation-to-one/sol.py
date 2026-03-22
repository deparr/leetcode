class Solution:
    def numSteps(self, s: str) -> int:
        digits = [int(c) for c in s]
        steps = 0
        while len(digits) > 1:
            steps += 1
            if digits[-1] == 0:
                digits.pop()
            else:
                i = len(digits) - 1
                while i >= 0 and digits[i] == 1:
                    steps += 1
                    i -= 1
                if i < 0:
                    break
                else:
                    digits[i] = 1
                    digits = digits[:i+1]

        return steps   

        
