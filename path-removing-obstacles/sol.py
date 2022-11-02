class Solution:
    def shortestPath(self, grid: List[List[int]], k: int) -> int:
        q = deque()
        vis = dict()
        # queue elements - (r, c, k left, steps so far)
        q.append((0, 0, k, 0))

        while q:
            next = q.popleft()
            if next[0] == len(grid)-1 and next[1] == len(grid[0])-1:
                return next[3]
            
            if next[0] < 0 or next[0] >= len(grid) or next[1] < 0 or next[1] >= len(grid[0]):
                continue
                
            if (next[0], next[1]) in vis:
                prev = vis[(next[0], next[1])]
                if prev[0] >= next[2]:
                    continue
            
            iswall = grid[next[0]][next[1]] == 1
            
            if iswall and next[2] == 0:
                continue
            
            
            nextk = next[2]
            if iswall:
                nextk -= 1
                
            vis[(next[0], next[1])] = (nextk, next[3])

            
            q.append((next[0]+1, next[1], nextk, next[3]+1))
            q.append((next[0], next[1]+1, nextk, next[3]+1))
            q.append((next[0], next[1]-1, nextk, next[3]+1))
            q.append((next[0]-1, next[1], nextk, next[3]+1))
                
                
        return -1
