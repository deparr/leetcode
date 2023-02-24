/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    alen, blen := 0, 0
    runner := headA

    for runner != nil {
        alen++
        runner = runner.Next
    }

    runner = headB
    for runner != nil {
        blen++
        runner = runner.Next
    }

    startdist := blen - alen
    var short *ListNode = headA
    runner = headB
    if alen > blen {
        startdist = alen - blen
        short = headB
        runner = headA
    }

    for i := 0; i < startdist; i++ {
        runner = runner.Next
    }

    for runner != nil {
        if runner == short {
            return runner
        }

        runner = runner.Next
        short = short.Next
    }
    return nil
}

