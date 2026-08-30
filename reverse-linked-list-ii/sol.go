/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var lag *ListNode
    cur := head
    fast := head.Next
    for i := 1; i < left; i++ {
        lag = cur
        cur = cur.Next
        fast = fast.Next
    }
    left_prev := lag
    left_start := cur

    for i := 0; i < right - left + 1; i++ {
        cur.Next = lag
        lag = cur
        cur = fast
        if fast != nil {
            fast = fast.Next
        } else {
            break
        }

    }

    if left_prev != nil {
        left_prev.Next = lag
    }
    left_start.Next = cur

    if left_prev == nil {
        return lag
    }

    return head
}
