# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        return self.sortedToBstHelper(nums, 0, len(nums) - 1,)
    
    def sortedToBstHelper(self, nums: List[int], lo: int, hi: int) -> Optional[TreeNode]:
        if lo > hi:
            return None
        if lo == hi:
            return TreeNode(nums[lo])
        mid = lo + (hi - lo) // 2
        root = TreeNode(nums[mid])
        root.left = self.sortedToBstHelper(nums, lo, mid - 1)
        root.right = self.sortedToBstHelper(nums, mid + 1, hi)
        return root
