package leetcode

import "sort"

// https://leetcode.cn/problems/h-index/?envType=study-plan-v2&envId=top-interview-150

func hIndex(citations []int) (h int) {
	// 首先我们可以将初始的 H\text{H}H 指数 hhh 设为 000，然后将引用次数排序，并且对排序后的数组从大到小遍历。
	// 根据 H\text{H}H 指数的定义，如果当前 H\text{H}H 指数为 hhh 并且在遍历过程中找到当前值 citations[i]>h\textit{citations}[i] > hcitations[i]>h，则说明我们找到了一篇被引用了至少 h+1h+1h+1 次的论文，所以将现有的 hhh 值加 111。继续遍历直到 hhh 无法继续增大。最后返回 hhh 作为最终答案。

	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}
