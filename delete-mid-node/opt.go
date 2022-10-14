// not quite the opt solution I don't think

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    fast, slow := head, head
    if head.Next == nil {
        head = nil
        return head
    }
    for true {
        fast = fast.Next.Next
        if fast == nil || fast.Next == nil {
            break
        }
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return head
}
