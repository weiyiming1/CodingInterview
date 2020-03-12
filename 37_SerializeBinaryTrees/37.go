package main

import (
	"CodingInterview/getOffer-Go/lib"
	"fmt"
	"strconv"
)

func Serialize(pRoot *lib.TreeNode, str []string) []string {
	if pRoot == nil {
		return append(str, "$")
	}

	str = append(str, strconv.Itoa(pRoot.Val))
	str = Serialize(pRoot.Left, str)
	str = Serialize(pRoot.Right, str)
	return str
}

func Deserialize(pRoot **lib.TreeNode, str *[]string) {
	temp := (*str)[0]
	*str = (*str)[1:]
	fmt.Println(str)
	if num, err := strconv.Atoi(temp); err == nil {
		*pRoot = new(lib.TreeNode)
		(*pRoot).Val = num
		(*pRoot).Left = nil
		(*pRoot).Right = nil

		Deserialize(&((*pRoot).Left), str)
		Deserialize(&((*pRoot).Right), str)
	}
}
