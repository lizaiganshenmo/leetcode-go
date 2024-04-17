package leetcode

// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
func maxVowels(s string, k int) int {
	// 思路 : 滑动窗口
	var (
		set = map[byte]struct{}{
			'a': {},
			'e': {},
			'i': {},
			'o': {},
			'u': {},
		}
	)
	n := len(s)
	left, right := 0, 0
	var tmp, ans int
	for ; right < n; right++ {
		if _, ok := set[s[right]]; ok {
			tmp++
		}

		if right-left+1 == k {
			ans = max1456(ans, tmp)
			if _, ok := set[s[left]]; ok {
				tmp--
			}
			left++
		}

	}
	return ans

}

func max1456(a, b int) int {
	if a > b {
		return a
	}
	return b
}
