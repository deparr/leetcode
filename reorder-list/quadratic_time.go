/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reorderList(head *ListNode)  {
   curr, runner, runnerprev := head, head.Next, head

   for curr.Next != nil {
	   
	   for runner.Next != nil {
		   runnerprev = runner
		   runner = runner.Next
	   }

	   runnerprev.Next = nil
	   runner.Next = curr.Next
	   curr.Next = runner

	   curr = runner.Next
	   if curr == nil {
		   break
	   }
	   runner = curr.Next
	   runnerprev = curr
   }
}
