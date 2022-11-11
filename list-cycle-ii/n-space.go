/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	cache := map[*ListNode]struct{}{}

	for head != nil {
		_, prs := cache[head]
		if prs {
			break
		}
		cache[head] = struct{}{}
		head = head.Next
	}
	return head
}