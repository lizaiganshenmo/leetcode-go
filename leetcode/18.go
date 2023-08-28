package leetcode

import "sort"

// 通用
// func fourSum(nums []int, target int) [][]int {
// 	n := len(nums)

// 	var ans [][]int
// 	if n < 4 {
// 		return ans
// 	}

// 	sort.Ints(nums)
// 	for j := 0; j < n-3; j++ {
// 		if j > 0 && nums[j] == nums[j-1] { //去重
// 			continue
// 		}
// 		for i := j + 1; i < n-2; i++ {
// 			if i > j+1 && nums[i] == nums[i-1] { //去重
// 				continue
// 			}
// 			left, right := i+1, n-1
// 			for left < right {
// 				t := nums[left] + nums[right] + nums[i] + nums[j]
// 				if t == target {
// 					ans = append(ans, []int{nums[j], nums[i], nums[left], nums[right]})
// 					left++
// 					right--
// 					for left < right && nums[left] == nums[left-1] {
// 						left++

// 					}
// 					for left < right && nums[right] == nums[right+1] {
// 						right--
// 					}

// 				} else if t > target {
// 					right--
// 				} else {
// 					left++
// 				}

// 			}
// 		}

// 	}

// 	return ans

// }

func fourSum(nums []int, target int) [][]int {
	// 三数之和 四数之和  k数和 最终都可转化成 两数之和
	// ksum target   => (k-1)sum target-nums[i]
	return KSum(nums, target, 4)

}

func KSum(nums []int, target, k int) [][]int {
	n := len(nums)
	var ans [][]int
	if n < k {
		return ans
	}

	sort.Ints(nums)

	var getSum func([]int, int, int, int, int) // // 数组, target, k, 遍历起点, 当前和
	path := make([]int, 0, k)                  // 记录当前遍历路径

	var twoSum func([]int, int, int) // 数组, target, 遍历起点
	twoSum = func(nums []int, target, start int) {
		end := len(nums) - 1
		// e := end
		// s := start
		for start < end {
			t := nums[start] + nums[end]
			if t > target {
				// 大了
				end--
			} else if t < target {
				// 小了
				start++
			} else {
				// 写入
				old := path
				path = append(path, nums[start], nums[end])
				ans = append(ans, append([]int{}, path...))
				// 复位
				path = old
				// fmt.Printf("path is : %+v\n", path)
				for start < end && nums[start] == nums[start+1] {
					start++
				}
				for start < end && nums[end] == nums[end-1] {
					end--
				}
				start++
				end--
			}
		}

	}

	getSum = func(nums []int, target, k, start, sumNow int) {
		if k <= 1 {
			return
		}
		if k == 2 {
			// fmt.Println("target-tmpSum now: ", target-tmpSum)
			twoSum(nums, target-sumNow, start)
			return
		}

		for i := start; i <= n-k; i++ {
			if nums[i] > 0 && nums[i]+sumNow > target {
				// 剪枝
				break
			}
			if i > start && nums[i] == nums[i-1] {
				// 遍历过了
				continue
			}

			// oldSum := tmpSum
			oldPath := path

			path = append(path, nums[i])
			getSum(nums, target, k-1, i+1, sumNow+nums[i])

			// 复位
			path = oldPath

		}

	}

	getSum(nums, target, k, 0, 0)
	return ans

}
