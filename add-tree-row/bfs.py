# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def addOneRow(self, root: Optional[TreeNode], val: int, depth: int) -> Optional[TreeNode]:
        if depth == 1:
            return TreeNode(val, root)
        q = deque()
        q.append(root)
        curr_depth = 1
        while curr_depth != depth - 1:
            to_process = len(q)
            for i in range(0, to_process):
                node = q.popleft()
                if node.right:
                    q.append(node.right)
                if node.left:
                    q.append(node.left)
            curr_depth += 1
        while q:
            parent = q.pop()
            new_left = TreeNode(val)
            new_right = TreeNode(val)
            if parent.left:
                new_left.left = parent.left
            if parent.right:
                new_right.right = parent.right
            parent.left = new_left
            parent.right = new_right
        return root 
