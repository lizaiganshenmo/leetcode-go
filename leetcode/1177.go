package leetcode

// https://leetcode.cn/problems/can-make-palindrome-from-substring/

func canMakePaliQueries(s string, queries [][]int) []bool {
	// 思路 : 模拟统计子串各字母个数
	// 直接做会超时 case无法全部通过 29/31  前缀和数组记录字符个数
	preSum := make([][26]int, len(s)+1)
	for i, c := range s {
		preSum[i+1] = preSum[i]
		preSum[i+1][c-'a']++
	}

	ans := make([]bool, len(queries))
	for i, v := range queries {
		tmpArr := [26]int{}
		for j := 0; j < 26; j++ {
			tmpArr[j] = preSum[v[1]+1][j] - preSum[v[0]][j]
		}
		ans[i] = canMake(tmpArr[:], v[2])
	}

	return ans

}

func canMake(sum []int, k int) bool {
	cnt := 0 // 奇数 字符 个数
	for _, val := range sum {
		if val%2 != 0 {
			cnt++
		}
	}

	// 若剩下单数字符可以 小于k*2 则可以构造回文串
	return cnt/2 <= k
}
