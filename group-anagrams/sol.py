class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)
        for s in strs:
            chars = str(sorted(s))
            groups[chars].append(s)
        return list(groups.values())

