package leetcode

// https://leetcode.cn/problems/count-number-of-ways-to-place-houses/description/

const mod int = 1e9 + 7

var f = [1e4 + 1]int{1, 2}

func init() {
	for i := 2; i <= 1e4; i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}
}

func countHousePlacements(n int) int {
	return f[n] * f[n] % mod
}
