/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil {
        return nil
    }
    count := 0
    runner, node, prev := head.Next, head, head

    for runner != nil {
        count++
        if count >= n {
            prev = node
            node = node.Next
        }
        runner = runner.Next
    }

    fmt.Println(runner, node, prev)

    if node == head {
        head = head.Next
    } else if node.Next == nil {
        prev.Next = nil
    } else {
        prev.Next = node.Next
    }

    return head

}
