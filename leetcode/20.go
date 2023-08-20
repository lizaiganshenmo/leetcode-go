package leetcode

// https://leetcode.cn/problems/valid-parentheses/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 示例 1：

// 输入：s = "()"
// 输出：true
// 示例 2：

// 输入：s = "()[]{}"
// 输出：true
// 示例 3：

// 输入：s = "(]"
// 输出：false

// 提示：

// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成

var (
	m = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
)

func isValid(s string) bool {
	leftStack := []byte{}
	for i := range s {
		switch s[i] {
		case '{', '[', '(':
			leftStack = append(leftStack, s[i])
		case '}', ']', ')':
			if len(leftStack) == 0 {
				return false
			}
			t := leftStack[len(leftStack)-1]
			leftStack = leftStack[0 : len(leftStack)-1]
			if t != m[s[i]] {
				return false
			}

		}

	}

	if len(leftStack) > 0 {
		return false
	}

	return true

}
