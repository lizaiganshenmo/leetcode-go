package leetcode

// https://leetcode.cn/problems/k-concatenation-maximum-sum/description/

func kConcatenationMaxSum(arr []int, k int) int {
	// 思路 : 动态规划
	// 参考 : https://leetcode.cn/problems/k-concatenation-maximum-sum/solutions/70400/java-kadanesuan-fa-yu-jie-ti-si-lu-by-zdxiq125/
	// 遍历到arr[i] 时, t = max(t, ans+arr[i])
	// ans = max(t, ans)
	const mod = 1e9 + 7
	var sum int
	var ans, curMax int
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	num := min1191(2, k)
	for i := 0; i < num*n; i++ {
		curMax = max1191(curMax+arr[i%n], arr[i%n])
		ans = max1191(curMax, ans)
	}

	if sum > 0 {
		k -= 2
		for k > 0 {
			ans = (ans + sum) % mod
			k--
		}
	}

	return ans % mod

}

func max1191(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min1191(a, b int) int {
	if a < b {
		return a
	}
	return b
}
