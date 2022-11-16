# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        
        hassum = False
        def explore(node, currsum):
            nonlocal hassum
            if not node.left and not node.right:
                if node.val + currsum == targetSum:
                    hassum = True
                return
            
            if node.left:
                explore(node.left, currsum + node.val)
            if node.right:
                explore(node.right, currsum + node.val)
        
        if root:
            explore(root, 0)
        return hassum
