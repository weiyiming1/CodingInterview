package main

func MoreThanHalfNum(num *[]int) int {
	// 空指针或空数组
	if num == nil || len(*num) == 0 {
		return -1
	}
	res := (*num)[0]
	count := 1
	for _, val := range *num {
		if count == 0 {
			res = val
			count = 1
		} else if val == res {
			count++
		} else {
			count--
		}
	}

	if checKMoreThanHalf(num, res) != true {
		return -1
	}
	return res

}

func checKMoreThanHalf(num *[]int, target int) bool {
	count := 0
	for _, val := range *num {
		if val == target {
			count++
		}
	}
	isMoreThanHalf := false
	if count*2 >= len(*num) {
		isMoreThanHalf = true
	}
	return isMoreThanHalf

}
