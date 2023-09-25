package leetcode

// https://leetcode.cn/problems/reverse-linked-list-ii/description/?envType=study-plan-v2&envId=top-interview-150

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dum := &ListNode{Next: head}
	pre, cur := dum, head
	length := right - left + 1

	left--
	for left > 0 && cur != nil {
		pre = pre.Next
		cur = cur.Next
		left--
	}
	lNodePre, lNode := pre, cur

	var rNode, rNodeNext *ListNode
	for length > 0 && cur != nil {

		next := cur.Next
		rNode, rNodeNext = cur, next
		cur.Next = pre
		pre = cur
		cur = next
		length--

	}

	lNode.Next = rNodeNext
	lNodePre.Next = rNode

	return dum.Next

}
