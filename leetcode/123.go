package leetcode

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/?envType=study-plan-v2&envId=top-interview-150

// func maxProfit3(prices []int) int {
// 	// 思路 : 动态规划
// 	// dp[i][j] 表示 [i,j] 区间内 第j天必卖股票获得的最大收益
// 	// dp[i][j] = max(dp[i][j1](这天卖股票最大收益) + prices[j]-[j1+1])(两次买卖), prices[j]-prices[i](一次买卖))
// 	const (
// 		DealMaxNum = 2 // 最大成交数
// 	)
// 	type SaleProgress struct {
// 		Profit  int // 收益
// 		DealNum int // 成交次数
// 	}

// 	var ans int
// 	n := len(prices)
// 	dp := make([][]*SaleProgress, n)
// 	for i := range dp {
// 		dp[i] = make([]*SaleProgress, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			sp := &SaleProgress{
// 				Profit:  prices[j] - prices[i], // 一次买卖
// 				DealNum: 1,
// 			}
// 			dp[i][j] = sp
// 			for t := i + 1; t < j; t++ {
// 				if dp[i][t].DealNum >= DealMaxNum {
// 					continue
// 				}
// 				tmpProfit := prices[j] - prices[t+1] + dp[i][t].Profit
// 				if tmpProfit > dp[i][j].Profit {
// 					dp[i][j].Profit = tmpProfit
// 					dp[i][j].DealNum++
// 				}

// 			}
// 			if dp[i][j].Profit > ans {
// 				ans = dp[i][j].Profit
// 			}
// 			fmt.Printf("dp[i][j] is: (i,j): (%d,%d) , val: %+v\n", i, j, dp[i][j])
// 		}

// 	}

// 	return ans

// }

func maxProfit3(prices []int) int {
	dp1 := make([]int, len(prices)) // dp1[i]表示截止到第i天，所能获得的最大利润
	dp2 := make([]int, len(prices)) // dp2[i]表示从第i天开始到最后一天，能获得的最大利润
	a := prices[0]                  // 到第i天股票的最低价
	b := prices[len(prices)-1]      // 从第i天开始股票的最高价

	// 两次遍历合并为一次
	for i := 1; i < len(prices); i++ {
		dp1[i] = max123(dp1[i-1], prices[i]-a)
		a = min123(a, prices[i])

		dp2[len(prices)-1-i] = max123(dp2[len(prices)-i], b-prices[len(prices)-1-i])
		b = max123(b, prices[len(prices)-1-i])
	}

	max := 0
	for i := 0; i < len(prices); i++ {
		max = max123(max, dp1[i]+dp2[i])
	}
	return max
}

func max123(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min123(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 作者：鄙人张麻子
// 链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/solutions/1420733/by-nervous-i2oentgen7rk-3fq7/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
