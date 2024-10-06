package leetcode

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/maximum-gap/description/

func maximumGap(nums []int) int {
	// O(nlogn)时间复杂度
	sort.Ints(nums)

	var ans int
	for i := 1; i < len(nums); i++ {
		t := nums[i] - nums[i-1]
		if t > ans {
			ans = t
		}

	}
	return ans

}

func maximumGap1(nums []int) int {
	// 思路: 桶排序 O(n)
	// 参考: https://leetcode.cn/problems/maximum-gap/solutions/2772242/xiang-xi-jie-shi-on-fen-tong-fa-pythonja-mi5s/

	// s1 查找数组中最大 最小值, 确定边界
	min, max := math.MaxInt, math.MinInt
	for _, val := range nums {
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	if max-min <= 1 {
		return max - min
	}

	//
	diff := max - min
	num := len(nums) - 1 // n个数 分成n-1段
	avg := diff / num
	if diff%num != 0 {
		avg++ // 向上取整
	}

	type pair struct{ min, max int } // 记录每个桶中 最大值 最小值
	buckets := make([]pair, (max-min)/avg+1)
	for i := range buckets {
		buckets[i] = pair{math.MaxInt, math.MinInt}
	}
	for _, x := range nums {
		b := &buckets[(x-min)/avg] // 分段看x 应该在哪个桶内
		b.min = min164(b.min, x)   // 维护桶内元素的最小值和最大值
		b.max = max164(b.max, x)
	}

	var ans int
	preMax := math.MaxInt
	for _, b := range buckets {
		if b.min != math.MaxInt { // 非空桶
			// 桶内最小值，减去上一个非空桶的最大值
			ans = max164(ans, b.min-preMax)
			preMax = b.max
		}
	}
	return ans

}

func min164(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max164(a, b int) int {
	if a > b {
		return a
	}
	return b
}
