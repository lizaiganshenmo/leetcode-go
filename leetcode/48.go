package leetcode

// https://leetcode.cn/problems/rotate-image/?envType=study-plan-v2&envId=top-interview-150

func rotate1(matrix [][]int) {
	// 思路:  图像顺时针旋转 90 度。 =》 1、上下对换、 2、主对角线对换

	n := len(matrix)
	// s1 上下对换
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// s2 主对角线对换
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
