package leetcode

// https://leetcode.cn/problems/longest-well-performing-interval/description/
func longestWPI(hours []int) (ans int) {
	n := len(hours)
	s := make([]int, n+1) // 前缀和
	st := []int{0}        // s[0]
	for j, h := range hours {
		j++
		s[j] = s[j-1]
		if h > 8 {
			s[j]++
		} else {
			s[j]--
		}
		if s[j] < s[st[len(st)-1]] {
			st = append(st, j) // 感兴趣的 j
		}
	}
	for i := n; i > 0; i-- {
		for len(st) > 0 && s[i] > s[st[len(st)-1]] {
			ans = max1124(ans, i-st[len(st)-1]) // [栈顶,i) 可能是最长子数组
			st = st[:len(st)-1]
		}
	}
	return
}

func max1124(a, b int) int {
	if b > a {
		return b
	}
	return a
}
