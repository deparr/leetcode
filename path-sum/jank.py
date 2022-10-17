# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        found = False
        def helper(node: TreeNode, _sum: int) -> bool:
            nonlocal found
            if not node:
                return True

            left, right = False, False
            left = helper(node.left, _sum + node.val)
            if found:
                return False
            right = helper(node.right, _sum + node.val)
            if found:
                return False

            if left and right:
                found = _sum + node.val == targetSum
            return False
        helper(root, 0)
        return found
