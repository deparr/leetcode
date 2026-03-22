/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    res := head
    carry := 0
    for l1 != nil || l2 != nil || carry != 0{
        n1, n2 := 0, 0
        if l1 != nil {
            n1 = l1.Val
        }
        if l2 != nil {
            n2 = l2.Val
        }

        sum := n1 + n2 + carry
        carry = sum / 10
        sum %= 10
        res.Val = sum

        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }

        if  l1 != nil || l2 != nil || carry > 0 {
            res.Next = &ListNode{}
        }
        res = res.Next
    }

    return head
}
