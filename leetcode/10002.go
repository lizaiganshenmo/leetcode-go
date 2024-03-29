package leetcode

import (
	"math/rand"
	"sort"
)

// https://blog.csdn.net/qq_45043813/article/details/126430933

// Card 扑克牌
type Card struct {
	Color int
	Num   int
}

// IsTongHuaShun 随机抽取num张牌，判断是否为同花顺
func IsTongHuaShun(cards []Card, num int) bool {
	n := len(cards)
	if n == 0 || n <= num || num <= 0 {
		return false
	}

	// 随机抽取num张
	set := make(map[int]struct{}, num)
	for len(set) < num {
		t := rand.Intn(n + 1)
		set[t] = struct{}{}
	}

	randList := make([]Card, 0, num)
	for k := range set {
		randList = append(randList, cards[k])
	}

	// 判断是否是同花顺
	sort.Slice(randList, func(i, j int) bool {
		return randList[i].Num < randList[j].Num
	})

	for i := 1; i < num; i++ {
		if randList[i].Color != randList[i-1].Color || randList[i].Num != randList[i-1].Num+1 {
			return false
		}
	}

	return true

}
