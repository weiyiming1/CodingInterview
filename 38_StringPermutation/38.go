package main

import "fmt"

func Permutation(pStr *[]string) {
	if pStr == nil {
		return
	}

	permutationDetail(pStr, 0)
}

func permutationDetail(pStr *[]string, beginIdx int) {
	if beginIdx == len(*pStr)-1 {
		fmt.Println(*pStr)
	} else {
		for curIdx := beginIdx; curIdx <= len(*pStr)-1; curIdx++ {
			// 交换，为了产生新的排序
			temp := (*pStr)[curIdx]
			(*pStr)[curIdx] = (*pStr)[beginIdx]
			(*pStr)[beginIdx] = temp

			permutationDetail(pStr, beginIdx+1)

			// 交换回去，为了循环结束打印当前排序
			temp = (*pStr)[curIdx]
			(*pStr)[curIdx] = (*pStr)[beginIdx]
			(*pStr)[beginIdx] = temp
		}
	}
}
