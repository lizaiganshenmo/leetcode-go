package leetcode

// https://leetcode.cn/problems/construct-string-with-repeat-limit/description/

func repeatLimitedString(s string, repeatLimit int) string {
	// 思路 :贪心
	// 优先使用z ->a
	n := len(s)
	if n == 0 {
		return ""
	}

	chNums := make([]int, 26)
	for i := range s {
		chNums[s[i]-'a']++
	}

	chBytes := make([]byte, n)
	idx := 0
	for i := 25; i >= 0; i-- {
		for chNums[i] > 0 {
			chBytes[idx] = 'a' + byte(i)
			chNums[i]--
			idx++
		}
	}

	// fmt.Printf("ch Bytes:%+v\n", chBytes)
	curCnt := 1
	for i := 1; i < n; i++ {
		if chBytes[i] == chBytes[i-1] {
			curCnt++
		} else {
			curCnt = 1
		}
		if curCnt > repeatLimit {
			ok := findDifchAndSwap(chBytes, i+1, i)
			if !ok {
				chBytes = chBytes[:i]
				break
			}
			curCnt = 1

		}
	}
	return string(chBytes)

}

// 返回找到的下标
func findDifchAndSwap(chs []byte, start int, changeIdx int) bool {
	n := len(chs)
	if start >= n {
		return false
	}

	for i := start; i < n; i++ {
		if chs[i] != chs[changeIdx] {
			chs[i], chs[changeIdx] = chs[changeIdx], chs[i]
			return true
		}
	}
	return false

}
