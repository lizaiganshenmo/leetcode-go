package leetcode

import "fmt"

// https://leetcode.cn/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150

func convert(s string, numRows int) string {
	n := len(s)
	if n == 1 {
		return s
	}

	// var cols int
	// if numRows > 1 {
	// 	cols = n / 2
	// 	if n%2 != 0 {
	// 		cols++
	// 	}
	// } else {
	// 	cols = n
	// }
	cols := n

	arr := make([][]byte, numRows)
	for i := range arr {
		arr[i] = make([]byte, cols)
	}

	fmt.Println("numRows cols:", numRows, cols)
	i, j := 0, 0
	flag := 0 // flag 0 向下走  1 右上走
	if i == numRows-1 {
		flag = 1
	}
	for idx := 0; idx < n; idx++ {
		fmt.Println("flag:", flag)
		fmt.Println("i,j:", i, j)
		arr[i][j] = s[idx]
		switch flag {
		case 0:
			// i++
			if i < numRows-1 {
				i++
			}

		case 1:
			j++
			// i--
			if i > 0 {
				i--
			}

		}
		if flag == 0 && (i == numRows-1 || numRows == 1) {
			flag = 1
		} else if flag == 1 && i == 0 && numRows != 1 {
			flag = 0
		}

	}

	// i,j = 0, 0
	// var sb strings.Builder
	fmt.Println(cols)
	bArr := make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		for j := 0; j < cols; j++ {
			if arr[i][j] != 0 {
				// _ = sb.WriteByte(arr[i][j])
				bArr = append(bArr, arr[i][j])
			}
		}
	}
	// return sb.String()
	return string(bArr)

}
