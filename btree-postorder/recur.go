:x
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func postorderTraversal(root *TreeNode) []int {
   res := make([]int, 0, 10)
   var post func(*TreeNode)
   post = func(node *TreeNode) {

	   if node.Left != nil {
		   post(node.Left)
	   }

	   if node.Right != nil {
		   post(node.Right)
	   }
	   res = append(res, node.Val)
   }

   if root != nil {
	   post(root)
   }
   return res
}
