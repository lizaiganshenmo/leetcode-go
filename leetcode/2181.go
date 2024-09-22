package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	// 思路 : 遍历模拟
	ans := &ListNode{}
	dummy := ans
	t := 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			node := &ListNode{Val: t}
			ans.Next = node
			ans = ans.Next
			t = 0
		} else {
			t += cur.Val
		}
	}

	return dummy.Next
}
