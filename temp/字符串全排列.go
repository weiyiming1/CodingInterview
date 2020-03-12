package main

import "fmt"

func permutation(arrayChar []byte, start int) {
	// 字串串长度0或1
	if len(arrayChar) <= 1 {
		return
	}
	// 起始排序节点在字符串末尾
	if start == len(arrayChar)-1 {
		for i := 0; i < len(arrayChar); i++ {
			fmt.Printf("%c", arrayChar[i])
		}
		fmt.Println()
	} else {
		for i := start; i < len(arrayChar); i++ {
			swap(arrayChar, start, i) // 获取新排列方式
			permutation(arrayChar, start+1)
			swap(arrayChar, start, i) // 恢复此次循环开始之前的排列方式
		}
	}
}

func swap(array []byte, m, n int) {
	temp := array[m]
	array[m] = array[n]
	array[n] = temp
}

func main() {
	arrayChar := []byte("abc")
	permutation(arrayChar, 0)
}
