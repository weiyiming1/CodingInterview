package main

import "fmt"

func FindDuplicate(array []int, length int, duplication *int) bool {

	if array == nil || length <= 0 {
		fmt.Println("空指针或空数据")
		return false
	}

	for _, temp := range array {
		if temp < 0 || temp > length-1 {
			fmt.Println("输入非法元素不符合题意")
			return false
		}
	}

	//主要部分
	for index, num := range array {
		for index != num {
			if num == array[num] {
				*duplication = num
				return true
			}

			//交换
			temp := num
			array[index] = array[temp]
			array[temp] = temp
		}
	}
	return false
}
