package leetcode

// https://leetcode.cn/problems/happy-number/description/?envType=study-plan-v2&envId=top-interview-150

func isHappy(n int) bool {
	mp := make(map[int]struct{})
	for n != 1 {
		if _, ok := mp[n]; ok {
			return false
		}
		mp[n] = struct{}{}
		n = nextStep(n)

	}

	return true

}

func nextStep(n int) int {
	var ans int
	for n != 0 {
		t := n % 10
		ans += t * t
		n /= 10
	}

	return ans

}
