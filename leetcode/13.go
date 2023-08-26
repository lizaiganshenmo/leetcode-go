package leetcode

import "fmt"

// https://leetcode.cn/problems/roman-to-integer/description/?envType=study-plan-v2&envId=top-interview-150

var (
	ma = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
)

func romanToInt(s string) int {
	n := len(s)

	var ans int
	for i := 0; i < n; i++ {
		fmt.Println("ans nonw is: ", ans)
		if i < n-1 && ma[s[i]] < ma[s[i+1]] {
			ans -= ma[s[i]]
		} else {
			ans += ma[s[i]]
		}
	}

	return ans

}
