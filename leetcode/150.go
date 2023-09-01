package leetcode

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/?envType=study-plan-v2&envId=top-interview-150

const (
	empty_val = -201
)

type Stack struct {
	list []int
}

func (s *Stack) Push(val int) {
	s.list = append(s.list, val)
}

func (s *Stack) Pop() int {
	if len(s.list) == 0 {
		return empty_val
	}
	t := s.list[len(s.list)-1]
	s.list = s.list[0 : len(s.list)-1]
	return t
}

func evalRPN(tokens []string) int {
	stack := Stack{}
	for _, val := range tokens {
		num, err := strconv.Atoi(val)
		if err != nil {
			mid := Cal(&stack, val)
			fmt.Printf("mid is : %d\n", mid)
			stack.Push(mid)

		} else {
			stack.Push(num)
		}
	}
	return stack.Pop()

}

func Cal(stack *Stack, label string) int {
	right := stack.Pop()
	left := stack.Pop()
	if left == empty_val || right == empty_val {
		return 0
	}

	var ans int
	switch label {
	case "+":
		ans = left + right
	case "-":
		ans = left - right
	case "*":
		ans = left * right
	case "/":
		ans = left / right
	default:
	}

	return ans

}
