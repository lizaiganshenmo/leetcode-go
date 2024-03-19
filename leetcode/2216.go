package leetcode

// https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/description/

func minDeletion(nums []int) int {
	// 参考: https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/solutions/2535327/gong-shui-san-xie-zhi-ji-ben-zhi-de-ji-j-dk05/
	var cnt int
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if (i-cnt)%2 == 0 && nums[i] == nums[i+1] {
			cnt++
		}

	}

	// 判断新数组是否为奇数
	if (n-cnt)%2 != 0 {
		cnt++
	}

	return cnt

}
