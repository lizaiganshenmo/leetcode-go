package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/largest-number/

func largestNumber(nums []int) string {
	// 思路 : 贪心算法
	// 优先组装 9 8 开头的大数目,9 93 99 优先取数目最短的,其次 99 93 排序
	m := make(map[int][]string, 10)
	for _, v := range nums {
		tmpStr := strconv.FormatInt(int64(v), 10)
		idx := int(tmpStr[0] - '0')
		m[idx] = append(m[idx], tmpStr)
	}

	var sb strings.Builder
	for i := 9; i >= 0; i-- {
		arr, ok := m[i]
		if !ok {
			continue
		}
		sort.Slice(arr, func(i, j int) bool {
			// x+y >y +x 则 x “大于” y
			xy, _ := strconv.ParseInt(arr[i]+arr[j], 10, 64)
			yx, _ := strconv.ParseInt(arr[j]+arr[i], 10, 64)
			if xy > yx {
				return true
			}
			return false

		})

		for _, v := range arr {
			sb.WriteString(v)
		}
	}

	ans := sb.String()
	if ans[0] == '0' {
		return "0"
	}
	return ans

}
