class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        
        points.sort(key=lambda p: p[0])
        
        start, stop = points[0][0], points[0][1]
        bmax = 1
        print(points)
        
        for i, p in enumerate(points):
            if i == 0:
                continue
            if p[0] == start:
                if p[1] < stop:
                    stop = p[1]
                continue
            
            if p[0] <= stop:
                if p[1] < stop:
                    stop = p[1]
                start = p[0]
                continue
            
            bmax += 1
            start, stop = p[0], p[1]

        return bmax
