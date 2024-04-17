package leetcode

// https://leetcode.cn/problems/grumpy-bookstore-owner/description/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// 滑动窗口
	// 哪个窗口生气的人最多,则将该窗口置控制情绪
	const (
		angryTag = 1
	)
	n := len(customers)
	left, right := 0, 0
	var ans int
	var tmp, madMaxNum int
	for ; right < n; right++ {
		if grumpy[right] == angryTag {
			tmp += customers[right]
		} else {
			ans += customers[right]
		}

		if right-left+1 == minutes {
			madMaxNum = max1052(madMaxNum, tmp)
			if grumpy[left] == angryTag {
				tmp -= customers[left]
			}
			left++

		}

	}

	return ans + madMaxNum

}

func max1052(a, b int) int {
	if a > b {
		return a
	}
	return b
}
