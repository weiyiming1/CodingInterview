package main

/*
二叉搜索树：
	若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
	若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
	它的左、右子树也分别为二叉排序树
*/

func VerifySquenceOfBST(sequence []int, len int) bool {
	if sequence == nil || len <= 0 {
		return false
	}

	root := sequence[len-1]
	var i int
	//在二叉搜索树中左子树节点的值小于根节点的值
	for i = 0; i < len-1; i++ {
		if sequence[i] > root {
			break
		}
	}

	//在二叉搜索树中右子树节点的值大于根节点的值
	for j := i; j < len-1; j++ {
		if sequence[j] < root {
			return false
		}
	}

	//判断左子树是不是二叉搜索树
	left := true
	if i > 0 {
		left = VerifySquenceOfBST(sequence, i)
	}

	//判断右子树是不是二叉搜索树
	right := true
	if i < len-1 {
		right = VerifySquenceOfBST(sequence[i:], len-i-1)
	}
	return left && right
}
