package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	dum := &ListNode{Next: head}
	for cur := dum; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return dum.Next

}
