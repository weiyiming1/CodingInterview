package main

import (
	"CodingInterview/getOffer-Go/lib"
	"testing"
)

func TestZPrint(t *testing.T) {
	t8 := lib.TreeNode{8, nil, nil}
	t9 := lib.TreeNode{9, nil, nil}
	t10 := lib.TreeNode{10, nil, nil}
	t11 := lib.TreeNode{11, nil, nil}
	t12 := lib.TreeNode{12, nil, nil}
	t13 := lib.TreeNode{13, nil, nil}
	t14 := lib.TreeNode{14, nil, nil}
	t15 := lib.TreeNode{15, nil, nil}
	t4 := lib.TreeNode{4, &t8, &t9}
	t5 := lib.TreeNode{5, &t10, &t11}
	t6 := lib.TreeNode{6, &t12, &t13}
	t7 := lib.TreeNode{7, &t14, &t15}
	t2 := lib.TreeNode{2, &t4, &t5}
	t3 := lib.TreeNode{3, &t6, &t7}
	t1 := lib.TreeNode{1, &t2, &t3}
	ZPrint(&t1)
}
