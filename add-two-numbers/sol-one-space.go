/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    res_head := l1
    res := l1
    for l1 != nil && l2 != nil {
        sum := l1.Val + l2.Val + carry
        carry = sum / 10
        l1.Val = sum % 10
        res = l1
        l1 = l1.Next
        l2 = l2.Next
    }

    var left_over *ListNode
    if l1 == nil && l2 == nil {
        if carry > 0 {
            res.Next = &ListNode{ Val: carry }
        }
    } else if l1 == nil {
        res.Next = l2
        left_over = l2
    } else {
        left_over = l1
    }

    for left_over != nil {
        sum := left_over.Val + carry
        carry = sum / 10
        left_over.Val = sum % 10
        if left_over.Next == nil {
            if carry > 0 {
                left_over.Next = &ListNode{ Val: carry }
                left_over = left_over.Next
            }
        }
        left_over = left_over.Next
    }

    return res_head
}
