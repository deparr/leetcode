/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    sz, idx := 0, 0
    tmp_head := head
    for tmp_head != nil {
        sz++
        tmp_head = tmp_head.Next
    }

    if sz == 1 {
        head = nil
    } else {
        tmp_head = head
        var prev *ListNode = nil
        for idx != sz / 2 {
            idx++
            prev = tmp_head
            tmp_head = tmp_head.Next
        }

        prev.Next = tmp_head.Next
    }

    return head
}
