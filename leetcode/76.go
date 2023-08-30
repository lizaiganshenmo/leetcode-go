package leetcode

func minWindow(s string, t string) string {
	// 思路： 滑动窗口
	n := len(s)
	if n == 1 {
		if s == t {
			return s
		} else {
			return ""
		}
	}

	need := make(map[byte]int)
	for i := range t {
		if val, ok := need[t[i]]; ok {
			need[t[i]] = val + 1
		} else {
			need[t[i]] = 1
		}
	}

	left, right := 0, 0
	flagNum := 0 // 记录窗口中匹配字符数目
	ansLen := n + 1
	var ans string
	m := len(t)
	for right < n {
		if val, ok := need[s[right]]; ok {
			need[s[right]] = val - 1
			if val > 0 {
				flagNum++
			}

		}

		if flagNum == m {
			// fmt.Printf("left now : %d , right now : %d\n", left, right)
			// 当前窗口已经满足元素
			winSize := right - left + 1
			if ansLen >= winSize {
				ansLen = winSize
				ans = s[left : right+1]
			}
			// fmt.Printf("need now :  %+v\n", need)
			// fmt.Printf("ans now :  %+v\n", ans)
			// 窗口左边界右移
			for left < right {
				if val, ok := need[s[left]]; ok {
					need[s[left]] = val + 1
					// fmt.Printf("need now :  %+v\n", need)
					// fmt.Printf("left now : %d , right now : %d  val : %+v\n", left, right, s[left:right+1])
					if val+1 > 0 {
						left++
						flagNum--
						break
					}
				}
				left++
				winSize := right - left + 1
				if ansLen >= winSize {
					ansLen = winSize
					ans = s[left : right+1]
				}
				// fmt.Printf("ans111 now :  %+v\n", ans)

			}

		}
		right++
	}

	return ans

}
