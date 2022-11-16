class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        deplist = {}
        for p in prerequisites:
            if p[1] not in deplist:
                deplist[p[1]] = [p[0]]
            else:
                deplist[p[1]].append(p[0])

        inlist = [0] * numCourses

        for node in deplist:
            for req in deplist[node]:
                inlist[req] += 1

        q = deque()
        for i, indeg in enumerate(inlist):
            if indeg == 0:
                q.append(i)

        res = []

        while q:
            node = q.popleft()
            res.append(node)
            if len(res) > numCourses:
                res.clear()
                break
            if node in deplist:
                for dep in deplist[node]:
                    inlist[dep] -= 1
                    if inlist[dep] == 0:
                        q.append(dep)
        if len(res) < numCourses:
            res.clear()
        return res

