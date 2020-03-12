package main

import "fmt"

func FindNumberAppearingOnce(num []int, res *int) {
	if len(num) == 0 {
		return
	}
	bitSum := make([]int, 32)
	for i := 0; i < len(num); i++ {
		bitMask := 1
		for j := 31; j >= 0; j-- {
			bit := num[i] & bitMask
			if bit != 0 {
				bitSum[j] += 1
				bitMask = bitMask << 1
			}
		}
		fmt.Println(bitMask)
	}

	*res = 0
	for i := 0; i < 32; i++ {
		*res = *res << 1
		*res += bitSum[i] % 3
	}
}
