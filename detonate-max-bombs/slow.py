class Solution:
    def maximumDetonation(self, bombs: List[List[int]]) -> int:
        res = []
        exploded = set()
        same_point = 0
        
        def explode(bomb, idx):
            nonlocal same_point

            exploded.add(idx)
            
            if len(exploded) == len(bombs):
                return
            
            for i, obomb in enumerate(bombs):
                if i not in exploded:
                    dist = (((obomb[0]-bomb[0])**2) + (obomb[1]-bomb[1])**2)**0.5
                    if dist <= bomb[2]:
                        explode(obomb, i)
                elif obomb[0] == bomb[0] and obomb[1] == bomb[1] and obomb != bomb:
                    same_point += 1

        for i, bomb in enumerate(bombs):
            if i not in exploded:
                exploded.clear()
                same_point = 0
                explode(bomb, i)
                res.append(len(exploded) + same_point)

        return max(res)
