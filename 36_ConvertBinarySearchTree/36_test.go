package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestConvert(t *testing.T) {
	t4 := lib.TreeNode{4, nil, nil}
	t8 := lib.TreeNode{8, nil, nil}
	t12 := lib.TreeNode{12, nil, nil}
	t16 := lib.TreeNode{16, nil, nil}
	t6 := lib.TreeNode{6, &t4, &t8}
	t14 := lib.TreeNode{14, &t12, &t16}
	t10 := lib.TreeNode{10, &t6, &t14}

	Convert(&t10)
	log.Println(t4.Val, t4.Left, t4.Right)
	log.Println(t6.Val, t6.Left, t6.Right)
	log.Println(t8.Val, t8.Left, t8.Right)
	log.Println(t10.Val, t10.Left, t10.Right)
	log.Println(t12.Val, t12.Left, t12.Right)
	log.Println(t14.Val, t14.Left, t14.Right)
	log.Println(t16.Val, t16.Left, t16.Right)

}
