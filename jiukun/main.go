package main

import (
	"fmt"
	"strings"
)

func main() {
	var set []string
	res := strings.Join(set, "\n")
	fmt.Println(res)
	err := fmt.Errorf(res)
	fmt.Println(err)
	fmt.Println(err == nil)
}

// rules
// 每次抽取 1 个筹码
// 如果抽取不满意，则放回重新抽取，直到抽到想要的筹码
// 对于取出的筹码，将相同面值的放在一起

// 总共有 kind 种筹码
// 九坤的目标为 nums
// 		nums[i] 表示第 i 堆筹码的数量

// 假设每种面值的筹码都有无限个，九坤总是遵循最优策略，是的他达成目标的操作次数最小

// 九坤达成目标的情况下，需要取出筹码次数的期望值

// [1,2] kind=4
// [0,0]

// 1 + 1 + 1/4 * 4/3 + 3/4 * 4/2 = 23 / 6
// [0,0]
// [1,0] 1
// [1,1] [2,0] 1
// [2,0] 1/4*(3/4*1+1/4*3/4*2+1/4*1/4*3/4*3...)

// 3/4 * (1/4^0+2/4^1+...+(n+1)(1/4)^n)

// (n+1)/(1/4)^n

// 3/4 * (1-1/4^n)/(1-1/4)

func chipGame(nums []int, kind int) float64 {
	return 0
}
