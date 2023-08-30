package leetcode

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/?envType=study-plan-v2&envId=top-interview-150

// func findSubstring(s string, words []string) []int {
// 	// todo : 此方法超时了，最终通过用例 : 151 / 179
// 	// s1 获取words的所有串联子串
// 	// 思路 : dfs 枚举
// 	n := len(words)
// 	canAdd := make(map[int]bool, n)
// 	for i := 0; i < n; i++ {
// 		canAdd[i] = true
// 	}
// 	// var subArr []string
// 	subArr := make(map[string]struct{})
// 	path := make([]string, 0, n)
// 	var dfs func(int) // 参数:当前已经入库几个
// 	dfs = func(num int) {
// 		// 递归终止条件
// 		if num == n {
// 			// subArr = append(subArr, strings.Join(path, ""))
// 			subArr[strings.Join(path, "")] = struct{}{}
// 			// path = make([]string, 0, n)
// 		}

// 		for i, can := range canAdd {
// 			if can {
// 				old := path
// 				path = append(path, words[i])
// 				canAdd[i] = false // 此次选了
// 				dfs(num + 1)

// 				// 复位
// 				path = old
// 				canAdd[i] = true
// 			}
// 		}

// 	}
// 	dfs(0)

// 	fmt.Printf("subArr is : %+v\n", subArr)
// 	// s2 处理subArr元素在s中下标
// 	var ans []int
// 	left, right := 0, n*len(words[0])
// 	sLen := len(s)
// 	for right <= sLen {
// 		t := s[left:right]
// 		if _, ok := subArr[t]; ok {
// 			ans = append(ans, left)
// 		}
// 		left++
// 		right++
// 	}

// 	return ans

// }

func findSubstring(s string, words []string) []int {
	// 思路 : 滑动窗口 比较窗口和words单词次数diff，一样则添加idx
	var ans []int
	n := len(words)
	sLen := len(s)
	step := len(words[0])
	left, right := 0, n*len(words[0])
	m := make(map[string]int, n)
	for _, v := range words {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}

	}

	for right <= sLen {
		tmpMap := make(map[string]int, n)
		for i := left; i < right; i += step {
			t := s[i : i+step]
			if val, ok := tmpMap[t]; ok {
				tmpMap[t] = val + 1
			} else {
				tmpMap[t] = 1
			}

		}
		if isSameMap(m, tmpMap) {
			ans = append(ans, left)
		}

		left++
		right++
	}

	return ans

}

func isSameMap(src, dst map[string]int) bool {
	n, m := len(src), len(dst)
	if n != m {
		return false
	}
	for i, v := range src {
		if val, ok := dst[i]; ok {
			if v != val {
				return false
			}
		} else {
			return false
		}

	}

	return true
}
