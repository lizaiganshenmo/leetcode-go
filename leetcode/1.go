package leetcode

// https://leetcode.cn/problems/two-sum/?envType=study-plan-v2&envId=top-interview-150

func twoSum1(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
