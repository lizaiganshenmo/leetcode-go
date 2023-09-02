package leetcode

// https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-interview-150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{Val: 0}
	if l1 != nil {
		pre.Next = l1
	} else {
		pre.Next = l2
	}

	carry := 0
	// var prePre *ListNode
	ans := pre
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + carry
		carry = tmp / 10
		node := &ListNode{Val: tmp % 10}
		pre.Next = node
		pre = node
		l1 = l1.Next
		l2 = l2.Next

	}

	for l1 != nil {
		tmp := l1.Val + carry
		carry = tmp / 10
		node := &ListNode{Val: tmp % 10}
		pre.Next = node
		pre = node
		l1 = l1.Next
	}
	for l2 != nil {
		tmp := l2.Val + carry
		carry = tmp / 10
		node := &ListNode{Val: tmp % 10}
		pre.Next = node
		pre = node
		l2 = l2.Next
	}

	if carry > 0 {
		node := &ListNode{Val: carry}
		pre.Next = node
	}

	return ans.Next

}
