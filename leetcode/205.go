package leetcode

// https://leetcode.cn/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 205. 同构字符串

// 给定两个字符串 s 和 t ，判断它们是否是同构的。

// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

//示例 1:

// 输入：s = "egg", t = "add"
// 输出：true
// 示例 2：

// 输入：s = "foo", t = "bar"
// 输出：false
// 示例 3：

// 输入：s = "paper", t = "title"
// 输出：true

func isIsomorphic(s string, t string) bool {
	return isSingle(s, t) && isSingle(t, s)

}

func isSingle(s, t string) bool {
	// 思路 维护一个map[byte]byte 表示s->t中字符的映射关系
	m := make(map[byte]byte)
	for i := range s {
		if val, ok := m[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}

	return true
}
