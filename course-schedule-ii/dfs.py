class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        prepost = [[-1, -1, i] for i in range(numCourses)]

        deplist = {}
        for p in prerequisites:
            if p[1] not in deplist:
                deplist[p[1]] = {p[0]}
            else:
                deplist[p[1]].add(p[0])

        currorder = 1
        cycle = False
        def explore(node, target):
            nonlocal currorder
            nonlocal cycle
            
            vis.add(node)
            if node in currpathvis:
                cycle = True
                return

            prepost[node][0] = currorder
            currorder += 1
            if node in deplist:
                currpathvis.add(node)
                for dep in deplist[node]:
                    if dep not in vis:
                        explore(dep, node)
                    elif prepost[dep][1] == -1:
                        cycle = True
                        return
                    if cycle:
                        return

            prepost[node][1] = currorder
            currorder += 1


        vis = set()
        currpathvis = set()
        for i in range(numCourses):
            if i not in vis:
                explore(i, i)
            currpathvis.clear()
            if cycle:
                return []

        res = []
        prepost.sort(key=lambda x: x[1], reverse=True)
        for node in prepost:
            res.append(node[2])
        return res
