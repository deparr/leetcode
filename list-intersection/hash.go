/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    set := map[*ListNode]struct {} {}

    for headA != nil {
        set[headA] = struct {} {}
        headA = headA.Next
    }

    for headB != nil {
        _, prs := set[headB]
        if prs {
            return headB
        }
        headB = headB.Next
    }

    return nil
}

