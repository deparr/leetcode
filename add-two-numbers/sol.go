func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	tmp := res
	carry := 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}

		sum := n1 + n2 + carry
		if sum >= 10 {
			carry = sum / 10
			sum -= 10
		} else {
			carry = 0
		}
		fmt.Println(carry)
		tmp.Val = sum

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if (l1 != nil) || (l2 != nil) {
			tmp.Next = &ListNode{0, nil}
			tmp = tmp.Next
		} else if carry > 0 {
			tmp.Next = &ListNode{0, nil}
			tmp = tmp.Next
		}

	}
	return res
}
