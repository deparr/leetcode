# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        count = 0
        def walk(node) -> tuple[int, int]:
            nonlocal count
            if node is None:
                return (0, 0)
            
            sub_sum = node.val
            n = 1
            l_sum, l_n = walk(node.left)
            r_sum, r_n = walk(node.right)

            sub_sum += l_sum + r_sum
            n += l_n + r_n
            avg = sub_sum // n
            if avg == node.val:
                count += 1
            
            return sub_sum, n

        walk(root)
        return count
