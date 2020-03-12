package main

//运算规则：0&0=0; 0&1=0; 1&0=0; 1&1=1;
func BinaryCountOne(num int) int {
	count := 0
	for num != 0 {
		// 整数二进制位有多少个1，就可以进行多少次位与运算
		num = (num - 1) & num
		count++
	}
	return count
}
