package leetcode

// https://leetcode.cn/problems/merge-in-between-linked-lists/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 1669. 合并两个链表
// 给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

// 请你将 list1 中下标从 a 到 b 的全部节点都删除，并将list2 接在被删除节点的位置。

// 下图中蓝色边和节点展示了操作后的结果：

// 请你返回结果链表的头指针。

// 示例 1：

// 输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
// 输出：[0,1,2,1000000,1000001,1000002,5]
// 解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。

// 3 <= list1.length <= 104
// 1 <= a <= b < list1.length - 1
// 1 <= list2.length <= 104

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dum := &ListNode{}
	dum.Next = list1
	pre := dum
	next := dum
	removeNum := b - a + 2
	for a > 0 {
		pre = pre.Next
		next = next.Next
		a--
	}
	for removeNum > 0 {
		next = next.Next
		removeNum--
	}

	// 链表中去除下标a->b节点
	pre.Next = list2
	cur := list2
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = next
	return list1

}
