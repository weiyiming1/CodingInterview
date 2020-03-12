package main

import (
	"CodingInterview/getOffer-Go/lib"
	"testing"
)

func TestFindPath(t *testing.T) {
	t4 := lib.TreeNode{4, nil, nil}
	t7 := lib.TreeNode{7, nil, nil}
	t5 := lib.TreeNode{5, &t4, &t7}
	t12 := lib.TreeNode{12, nil, nil}
	t10 := lib.TreeNode{10, &t5, &t12}

	FindPath(&t10, 22)

}
