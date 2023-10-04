package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/greatest-sum-divisible-by-three/

// todo : 此题没完全Ac 40/42
func maxSumDivThree(nums []int) int {
	// 思路 : 贪心算法
	arr := make([][]int, 3)
	var ans int
	for _, v := range nums {
		key := v % 3
		if key == 0 {
			ans += v
		} else {
			arr[key] = append(arr[key], v)
		}

	}
	// fmt.Println("ans is: ", ans)

	for t := 1; t <= 2; t++ {
		sort.Slice(arr[t], func(i, j int) bool {
			return arr[t][i] > arr[t][j]
		})
	}

	n1, n2 := len(arr[1]), len(arr[2])
	idx1, idx2 := 0, 0
	for idx1 < n1 && idx2 < n2 {
		var t1, t2, t3 int
		t3 = arr[1][idx1] + arr[2][idx2]
		if idx1+2 < n1 {
			t1 = arr[1][idx1] + arr[1][idx1+1] + arr[1][idx1+2]
		}
		if idx2+2 < n2 {
			t2 = arr[2][idx2] + arr[2][idx2+1] + arr[2][idx2+2]
		}
		if t3 >= t1 && t3 >= t2 {
			ans += t3
			idx1++
			idx2++
		} else if t2 >= t1 && t2 >= t3 {
			ans += t2
			idx2 += 3

		} else {
			ans += t1
			idx1 += 3
		}

	}

	for idx2 < n2 && idx2+2 < n2 {
		ans += arr[2][idx2] + arr[2][idx2+1] + arr[2][idx2+2]
		idx2 += 3
	}
	for idx1 < n1 && idx1+2 < n1 {
		ans += arr[1][idx1] + arr[1][idx1+1] + arr[1][idx1+2]
		idx1 += 3
	}

	return ans

}
