package leetcode

// https://leetcode.cn/contest/weekly-contest-366/problems/divisible-and-non-divisible-sums-difference/

func differenceOfSums(n int, m int) int {
	tmp := m
	sum, jianSum := 0, 0
	for i := 1; i <= n; i++ {
		sum += i
		if i == m {
			jianSum += i
			m += tmp
		}
	}
	// fmt.Println("sum: ", sum)
	// fmt.Println("jianSum: ", jianSum)

	return sum - 2*jianSum

}
