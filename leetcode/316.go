package leetcode

// https://leetcode.cn/problems/remove-duplicate-letters/

func removeDuplicateLetters(s string) string {
	left := ['z' + 1]int{} // 相比创建一个长为 26 的数组，多开一点空间更方便
	for _, c := range s {
		left[c]++ // 统计每个字母的出现次数
	}
	ans := []rune{}
	inAns := ['z' + 1]bool{}
	for _, c := range s {
		left[c]--
		if inAns[c] { // ans 中不能有重复字母
			continue
		}
		for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
			// 如果 c < x，且右边还有 x，那么可以把 x 去掉，因为后面可以重新把 x 加到 ans 中
			x := ans[len(ans)-1]
			ans = ans[:len(ans)-1]
			inAns[x] = false // 标记 x 不在 ans 中
		}
		ans = append(ans, c) // 把 c 加到 ans 的末尾
		inAns[c] = true      // 标记 c 在 ans 中
	}
	return string(ans)
}

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/remove-duplicate-letters/solutions/2381483/gen-zhao-wo-guo-yi-bian-shi-li-2ni-jiu-m-zd6u/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
