# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        seen = set()
        return Solution.search(root, seen, k)

    def search(r: Optional[TreeNode], seen: set[int], k) -> bool:
        if r is None:
            return False
        if k - r.val in seen:
            return True
        seen.add(r.val)
        return Solution.search(r.left, seen, k) or Solution.search(r.right, seen, k)
