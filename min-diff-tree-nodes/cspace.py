# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDiffInBST(self, root: Optional[TreeNode]) -> int:
        _min = 10**5 + 1
        prev = None
        
        def explore(node):
            nonlocal prev
            nonlocal _min
            if not node:
                return
            
            explore(node.left)
            if prev and abs(node.val - prev.val) < _min:
                _min = abs(node.val - prev.val)
            prev = node
            explore(node.right)

        explore(root)
        
        return _min

