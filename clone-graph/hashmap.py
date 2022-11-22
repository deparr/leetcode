"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if not node:
            return None

        if not node.neighbors:
            return Node(1)

        newhead = Node(1)
        valtonode = defaultdict(Node)
        valtonode[1] = newhead
        vis = set()
        q = deque()
        q.append(node)
        currparent = None

        while q:
            _next = q.popleft()
            vis.add(_next.val)

            if _next.val not in valtonode:
                valtonode[_next.val] = Node(_next.val)

            for node in _next.neighbors:

                if node.val not in vis and node.val not in valtonode:
                    q.append(node)

                toadd = None
                if node.val not in valtonode:
                    toadd = Node(node.val)
                    valtonode[node.val] = toadd
                else:
                    toadd = valtonode[node.val]

                valtonode[_next.val].neighbors.append(toadd)

        return newhead

