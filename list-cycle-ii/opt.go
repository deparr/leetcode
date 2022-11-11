/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast, entry := head, head, head

	if fast == nil {
		return nil
	}

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			for slow != entry {
				slow = slow.Next
				entry = entry.Next
			}
			return slow
		}
	}

	return nil
}