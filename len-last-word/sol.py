def lengthOfLastWord(self, s: str) -> int:
	strs = s.rstrip().split(sep=' ')
	return len(strs[-1])
