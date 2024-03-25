package leetcode

// https://leetcode.cn/problems/fruit-into-baskets/description/

func totalFruit(fruits []int) int {
	// 思路: 滑动窗口
	const (
		fruitKinds = 2
	)
	n := len(fruits)
	if n <= fruitKinds {
		return n
	}

	left := 0
	mp := make(map[int]int) // 窗口内元素val 对应数目
	mp[fruits[0]] = 1
	// cnt := 0 // 窗口中水果种类数
	var ans int
	for right := 1; right < n; right++ {
		mp[fruits[right]]++
		// 窗口左边界右移
		for ; len(mp) > fruitKinds; left++ {
			mp[fruits[left]]--
			if mp[fruits[left]] == 0 {
				delete(mp, fruits[left])
			}

		}
		ans = max904(ans, right-left+1)
	}

	return ans

}

func max904(a, b int) int {
	if a > b {
		return a
	}
	return b
}
