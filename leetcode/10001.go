package leetcode

import (
	"slices"
	"sort"
	"strconv"
)

// https://blog.csdn.net/weixin_44823312/article/details/127836253
// https://blog.csdn.net/qq_36282995/article/details/126078742

// GetMaxNumLessN 小于N的最大数字
func GetMaxNumLessN(arr []int, num int) int {
	// 思路: dfs+贪心
	// {2,3,8}
	// 1111 情况: 高位没有更小的且高位 直接舍弃最高位
	// 9999 情况: 高位没有更大于的,直接最大8888  , 维护一个bool变量alreadyLess 保证已经有高位小于该对应位置,则后面地位可以直接最大的填充
	// 2883 情况: 前三个288 , 后面一位取2
	// 2882 情况: 前三个288,到第四位发现没有更小的,则第三位8降为3 ,最后一位变为8 2838
	const (
		invalidNum = -1
		notExist   = -2
	)
	if num <= 0 {
		return invalidNum
	}

	n := len(arr)
	if n == 0 {
		return 0
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	numStr := strconv.Itoa(num)
	path := make([]byte, 0, n)
	var dfs func(arr []int, s string, alreadyLess bool, idx int) // 参数:arr: 倒序int数组 s: 数字字符串 ,alreadyLess 当前是否已经存在高位小于num对应高位, idx:当前搜索下标
	dfs = func(arr []int, s string, alreadyLess bool, idx int) {
		// 递归终止条件
		if idx == len(s) {
			return
		}

		// 如果已经存在高位小于num对应高位
		if alreadyLess {
			path = append(path, byte(arr[0]+'0'))
			dfs(arr, s, alreadyLess, idx+1)
		} else {
			// 还不存在, 分别判断
			t := int(s[idx] - '0')
			val := invalidNum
			for i := 0; i <= len(arr)-1; i++ {
				if idx == len(s)-1 {
					if arr[i] < t {
						val = arr[i]
						break
					}
				} else {
					if arr[i] <= t {
						val = arr[i]
						break
					}
				}

			}
			if val == invalidNum {
				// 此时idx 走到最后一位, 且该位val等于num该位对应的数 如{2,3,8} 2882
				if idx == len(s)-1 {
					// 回溯
					i := len(path) - 1
					for ; i >= 0; i-- {
						j := slices.Index(arr, int(path[i]-'0'))
						if j >= 0 && j < len(arr)-1 {
							path[i] = byte(arr[j+1] + '0')
							path = path[:i+1]
							break
						}

					}
					if i < 0 {
						// 未回溯成功,说明需要直接少一位,path清空
						path = []byte{}
						i = 0
					}
					dfs(arr, s, true, i+1)
				} else {
					// 不存在小于等于该位的,例如{2,3,8} 1111 则直接舍弃该位
					alreadyLess = true
					dfs(arr, s, alreadyLess, idx+1)
				}
			} else {
				path = append(path, byte(val+'0'))
				dfs(arr, s, alreadyLess, idx+1)
			}

		}

	}

	dfs(arr, numStr, false, 0)

	if len(path) == 0 {
		return notExist
	}

	res, _ := strconv.Atoi(string(path))
	return res

}
