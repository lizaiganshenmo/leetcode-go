package leetcode

// https://leetcode.cn/problems/minimum-number-of-operations-to-make-array-empty/

func minOperations1(nums []int) int {
	// 贪心思想 : 尽量每次3个一组消除  4%3 == 1(2*2)=两次 情况  10%3==1 (2*3+4)=4次  if x % 3 == 1 cnt = x/3+1
	//                             5%3== 2(3+2) =两次   8%3== 2(2*3+2) = 4 次 if x % 3 == 1 cnt = x/3+1
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	cnt := 0
	for _, v := range m {
		if v == 1 {
			return -1
		}
		if v == 2 {
			cnt++

		} else {
			cnt += v / 3
			if v%3 != 0 {
				cnt++
			}
		}

	}

	return cnt

}
