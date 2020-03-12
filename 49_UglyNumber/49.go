package main

func GetUglyNumber(idx int) int {
	if idx <= 0 {
		return 0
	}
	uglyNum := make([]int, idx)
	uglyNum[0] = 1
	nextIdx := 1
	multi2 := 0
	multi3 := 0
	multi5 := 0

	for nextIdx < idx {
		min := getMin(uglyNum[multi2]*2, uglyNum[multi3]*3, uglyNum[multi5]*5)
		uglyNum[nextIdx] = min
		for uglyNum[multi2]*2 <= uglyNum[nextIdx] {
			multi2++
		}
		for uglyNum[multi3]*3 <= uglyNum[nextIdx] {
			multi3++
		}
		for uglyNum[multi5]*5 <= uglyNum[nextIdx] {
			multi5++
		}
		nextIdx++
	}
	return uglyNum[nextIdx-1]
}

func getMin(num1, num2, num3 int) int {
	var min int
	if num1 < num2 {
		min = num1
	} else {
		min = num2
	}
	if min < num3 {
		return min
	} else {
		return num3
	}
}
