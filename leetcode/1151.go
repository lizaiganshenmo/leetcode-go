package leetcode

// https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together/description/
func minSwaps(data []int) int {
	// 思路 : 滑动窗口
	// 维护一个 winSize = data数组中 1个数 的固定窗口 ,
	// 窗口右滑过程中，时刻记录窗口中 1的数目 tmp, 则最少步数应该是 winSize - max(tmp)

	n := len(data)
	winSize := 0
	for _, val := range data {
		if val == 1 {
			winSize++
		}
	}

	left, right := 0, winSize-1
	tmp := 0
	maxOneNum := 0
	for i := left; i <= right; i++ {
		if data[i] == 1 {
			tmp++
		}
	}
	for right < n {
		maxOneNum = max1151(maxOneNum, tmp)
		if left < n && data[left] == 1 {
			tmp--
		}
		left++
		right++
		if right < n && data[right] == 1 {
			tmp++
		}
	}

	return winSize - maxOneNum

}

func max1151(a, b int) int {
	if a > b {
		return a
	}
	return b
}
