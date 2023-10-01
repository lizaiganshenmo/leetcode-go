package leetcode

// https://leetcode.cn/contest/weekly-contest-365/problems/maximum-value-of-an-ordered-triplet-i/

func maximumTripletValue(nums []int) int64 {
	// 思路 : 动态规划
	// 找到以 i 为点的 右边的最小值和最大值 对应的下标
	//
	n := len(nums)
	if n == 3 {
		tmp := int64((nums[0] - nums[1]) * nums[2])
		if tmp < 0 {
			tmp = 0
		}
		return tmp
	}
	dpMax, dpLeft := make([]int, n), make([]int, n)
	dpMax[n-2] = n - 1

	for i := n - 3; i >= 0; i-- {
		if nums[i+1] > nums[dpMax[i+1]] {
			dpMax[i] = i + 1
		} else {
			dpMax[i] = dpMax[i+1]
		}

	}
	dpLeft[1] = 0
	for i := 2; i < n; i++ {
		if nums[i-1] > nums[dpLeft[i-1]] {
			dpLeft[i] = i - 1
		} else {
			dpLeft[i] = dpLeft[i-1]
		}

	}

	var maxNum int64
	for i := 1; i <= n-2; i++ {
		tmp := int64((nums[dpLeft[i]] - nums[i]) * nums[dpMax[i]])
		if tmp > maxNum {
			maxNum = tmp

		}

	}
	return int64(maxNum)

}
