package leetcode

// https://leetcode.cn/problems/moving-average-from-data-stream/

type MovingAverage struct {
	size        int
	left, right int
	sizeCount   int
	arr         []int
}

func Constructor3(size int) MovingAverage {
	return MovingAverage{
		size: size,
		arr:  make([]int, 0),
	}

}

func (this *MovingAverage) Next(val int) float64 {
	this.arr = append(this.arr, val)
	this.right++
	this.sizeCount += val
	if this.right-this.left > this.size {
		// fmt.Println("this.right", this.right, val)
		this.sizeCount -= this.arr[this.left]
		this.left++
	}
	cap := len(this.arr)
	if cap > this.size {
		cap = this.size
	}
	// fmt.Println("this.sizeCount", this.sizeCount, cap)
	return float64(this.sizeCount) / float64(cap)

}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
