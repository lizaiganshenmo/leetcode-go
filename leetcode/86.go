package leetcode

// https://leetcode.cn/problems/M99OJA/description/

func partition(s string) (ans [][]string) {
	// 思路: dfs
	// 判断当前截取子串是否是回文串,是的话继续dfs

	var path []string
	var dfs func(s string)
	dfs = func(s string) {
		n := len(s)
		if n == 0 && len(path) != 0 {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for i := 1; i <= n; i++ {
			if isHuiwenStr(s[0:i]) {
				old := path
				path = append(path, s[0:i])
				dfs(s[i:])

				// 复位
				path = old
			}
		}

	}

	dfs(s)

	return

}

func isHuiwenStr(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
	}
	l++
	r--

	return true
}
