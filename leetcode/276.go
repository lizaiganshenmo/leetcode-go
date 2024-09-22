package leetcode

// https://leetcode.cn/problems/paint-fence/description/?envType=problem-list-v2&envId=dynamic-programming&difficulty=MEDIUM
func numWays(n int, k int) int {

	// 参考: https://leetcode.cn/problems/paint-fence/?envType=problem-list-v2&envId=dynamic-programming&difficulty=MEDIUM

	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	// 到第 i 个有几种涂色的方式
	dp := make([]int, n)
	dp[0] = k
	dp[1] = k * k

	/*

	   n > 3 时，有 2 中互斥的状态，

	   1. 第 i 个栅栏使用与 i-1 同色，则有：
	       dpsame = dp[i-2] * (k-1) * 1

	       * k-1 - 第 i-1 栅栏可以选择方式的个数
	       * 乘以 1 表示，第 i 个栅栏因为要与第 i-1 栅栏相同

	   2. 第 i 个栅栏使用与 i-1 不同色，则有：
	       dpdiff = dp[i-1] * (k-1)

	       * k-1 - 因为不同色，所以一共有 k-1 种方式可以选择。

	*/

	for i := 2; i < n; i++ {

		dp[i] = (dp[i-2] * (k - 1) * 1) + (dp[i-1] * (k - 1))
	}

	return dp[n-1]
}
