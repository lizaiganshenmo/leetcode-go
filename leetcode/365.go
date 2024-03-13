package leetcode

// https://leetcode.cn/problems/water-and-jug-problem/description/

func canMeasureWater(x int, y int, target int) bool {
	// 如果两个x 、 y 的最大公约数, 可以被target整除 则可以
	if x+y < target {
		return false
	}
	t := getMaxYue(x, y)
	return target%t == 0

}

func getMaxYue(x, y int) int {
	t := x % y
	for t != 0 {
		x = y
		y = t
		t = x % y
	}
	return y
}
