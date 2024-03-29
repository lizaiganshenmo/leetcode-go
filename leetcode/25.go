package leetcode

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	isFinalHead := false
	// 需要记录翻转后的tail,即原来该部分链表的head
	tail := head
	cnt := 0
	curHead := head
	var ans, preTailNode *ListNode
	for cur := head; cur != nil; {
		cnt++
		if cnt == k {
			cnt = 0
			tmpNode := cur.Next
			cur.Next = nil // 断链
			root := reverseLinkList(curHead, preTailNode)
			// fmt.Printf("root is:%+v\n", root)

			if !isFinalHead {
				isFinalHead = true
				ans = root
			}
			preTailNode = curHead
			// fmt.Printf("preTailNode is:%+v\n", preTailNode)
			curHead = tmpNode
			cur = tmpNode
			tail.Next = tmpNode
			tail = cur
			continue
		}
		cur = cur.Next

	}

	return ans

}

// 翻转链表
func reverseLinkList(head, preTailNode *ListNode) *ListNode {
	dum := &ListNode{
		Next: head,
	}
	pre, cur := dum, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if preTailNode != nil {
		preTailNode.Next = pre
	}

	return pre

}
