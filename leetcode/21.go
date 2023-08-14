package leetcode

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/?envType=study-plan-v2&envId=huawei-2023-spring-sprint

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

// 示例：

// 输入：nums = [1,2,3,4]
// 输出：[1,3,2,4]
// 注：[3,1,2,4] 也是正确的答案之一。

// 提示：

// 0 <= nums.length <= 50000
// 0 <= nums[i] <= 10000

func exchange(nums []int) []int {
	n := len(nums)
	i, j := 0, n-1

	for i < j {
		for i < j && isJiShu(nums[i]) {
			i++
		}
		for i < j && !isJiShu(nums[j]) {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums

}

func isJiShu(i int) bool {
	if i%2 != 0 {
		return true
	}
	return false
}
