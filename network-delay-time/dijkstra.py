class Solution:
def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
    
    edges = {}
    edgecount = 0
    for t in times:
        if t[0] not in edges:
            edges[t[0]] = [[t[1], t[2]]]
            edgecount += 1
        else:
            edges[t[0]].append([t[1], t[2]])
            edgecount += 1
    
    if edgecount < n-1 or k not in edges:
        return -1
    
    dp = {}
    for i in range(1, n+1):
        dp[i] = [i, -1, -1]
        if i == k:
            dp[i][1] = 0
            

    q = deque()
    q.append(dp[k])
    while q:
        _next = q.popleft()
        if _next[0] in edges:
            for edge in edges[_next[0]]:
                if dp[edge[0]][1] == -1 or dp[edge[0]][1] > dp[_next[0]][1] + edge[1]:
                    dp[edge[0]][1] = dp[_next[0]][1] + edge[1]
                    dp[edge[0]][2] = _next[0]
                    q.append(dp[edge[0]])

    _max = 0
    for distprev in dp.values():
        if distprev[1] == -1:
            return -1
        if distprev[1] > _max:
            _max = distprev[1]
    return _max
