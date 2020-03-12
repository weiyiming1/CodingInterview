package main

func FindNumAppearOnce(data []int, num1, num2 *int) {
	if len(data) < 2 {
		return
	}
	xor := 0
	for _, val := range data {
		xor ^= val
	}

	idxOf1 := findFirstBit1(xor)

	*num1 = 0
	*num2 = 0
	for i := 0; i < len(data); i++ {
		if isBit1(data[i], idxOf1) {
			*num1 = *num1 ^ data[i]
		} else {
			*num2 = *num2 ^ data[i]
		}
	}
}

// 在整数num的二进制表示中找到最右边是1的位
func findFirstBit1(num int) int {
	idx := 0
	for (num&1) == 0 && idx < 64 {
		num = num >> 1 // 右移一位
		idx++
	}
	return idx
}

//判断num的二进制从右边数起的idx位是不是1
func isBit1(num, idx int) bool {
	num = num >> uint(idx)
	return num&1 == 1
}
