class Solution:
    def checkIfPangram(self, sentence: str) -> bool:
        if len(sentence) < 26:
            return False
        for let in "abcdefghijklmnopqrstuvwxyz":
            if let not in sentence:
                return False
        return True
