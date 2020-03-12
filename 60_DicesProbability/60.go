package main

import (
	"fmt"
	"math"
)

// 计算n个骰子某次投掷点数和为s的出现次数
func countNum(n, s int) int {
	//n个骰子点数之和范围在n到6n之间，否则数据不合法
	if s < n || s > 6*n {
		return 0
	}
	//当有一个骰子时，一次骰子点数为s(1 <= s <= 6)的次数当然是1
	if n == 1 {
		return 1
	} else {
		return countNum(n-1, s-1) + countNum(n-1, s-2) + countNum(n-1, s-3) +
			countNum(n-1, s-4) + countNum(n-1, s-5) + countNum(n-1, s-6)
	}
}

func PrintProbability(n int) {
	total := int(math.Pow(float64(6), float64(n)))
	for i := n; i <= 6*n; i++ {
		fmt.Printf("P(sum=%d) = %d/%d\n", i, countNum(n, i), total)
	}
}

/*
教材，解法一：

const maxValue  = 6

func PrintProbability(num int){
	if num < 1{
		return
	}

	maxSum := num*maxValue
	pProbabilities := make([]int, maxSum-num+1)
	getProbability(num, pProbabilities)
	total := int(math.Pow(float64(maxValue),float64(num)))
	for i:=num; i<=maxSum;i++{
		fmt.Printf("P(sum=%d) = %d/%d\n", i, pProbabilities[i-num], total);
	}
}

func getProbability(num int, pProbabilities []int){
	for i:=1; i <= maxValue; i++{
		probability(num, num, i, pProbabilities)
	}
}

func probability(original, current, sum int, pProbabilities []int){
	// 只剩下一个骰子
	if current == 1{
		pProbabilities[sum-original]++
	}else {
		for i:=1;i<=maxValue;i++{
			probability(original, current-1, i+sum, pProbabilities)
		}
	}
}*/
