package leetcode

// https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/
func minimumRecolors(blocks string, k int) int {
	// 滑动窗口
	const (
		BlackTag = 'B'
		WhiteTag = 'W'
	)

	n := len(blocks)
	if n < k {
		return 0
	}
	left, right := 0, 0
	var tmp int
	var ans = k
	for ; right < n; right++ {
		if blocks[right] == BlackTag {
			tmp++
		}
		if right-left+1 == k {
			ans = min2379(ans, k-tmp)
			if blocks[left] == BlackTag {
				tmp--
			}
			left++
		}

	}
	return ans

}

func min2379(a, b int) int {
	if a < b {
		return a
	}
	return b
}
