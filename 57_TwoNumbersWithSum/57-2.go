package main

import (
	"fmt"
)

func FindContinuousSequence(sum int) {
	if sum < 3 {
		return
	}

	small := 1 // small 指针初始化为1
	big := 2   //  big 指针初始化为2
	mid := (1 + sum) / 2
	curSum := small + big

	for small < mid {
		if curSum == sum {
			printContinuousSeq(small, big)
		}

		for curSum > sum && small < mid {
			curSum -= small
			small++
			if curSum == sum {
				printContinuousSeq(small, big)
			}
		}
		big++
		curSum += big
	}
}

func printContinuousSeq(small, big int) {
	for i := small; i <= big; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}
