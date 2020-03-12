package main

// 参考思路： https://blog.csdn.net/yi_Afly/article/details/52012593

func NumberBetween1AndN(n int) int {
	if n < 1 {
		return 0
	}
	count := 0
	base := 1
	round := n
	for round > 0 {
		weight := round % 10
		round /= 10
		count += round * base
		if weight == 1 {
			// round*base+former+1
			count += (n % base) + 1
		} else if weight > 1 {
			// rount*base+base
			count += base
		}
		// 更高位
		base *= 10
	}
	return count
}
