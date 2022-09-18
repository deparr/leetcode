def maxDepth(self, root: Optional[TreeNode]) -> int:
    max = 0
    def helper(node: TreeNode, count: int):
        nonlocal max
        if not node:
            if count > max:
                max = count
            return
        helper(node.left, count + 1)
        helper(node.right, count + 1)
        return
    helper(root, 0)
    return max