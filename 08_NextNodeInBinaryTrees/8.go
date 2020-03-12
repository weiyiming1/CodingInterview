package main

type BinaryTreeNode struct {
	Value  string
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Father *BinaryTreeNode
}

func GetNextNode(curNode *BinaryTreeNode) *BinaryTreeNode {
	if curNode == nil {
		return nil
	}

	//节点有右子树，下个节点就是右子树的左子节点
	if curNode.Right != nil {
		temp := curNode.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		return temp

	} else if curNode.Father.Left == curNode { //节点没有右子树，它是父节点的左子节点，下个节点就是父节点
		return curNode.Father
	} else {
		//节点没有右子树，它是父节点的右子节点， 则沿着父指针向上遍历，
		// 直到发现一个节点位于它父节点的左侧，那么这个父节点就是所找的
		temp := curNode.Father
		for temp.Father != nil {
			if temp.Father.Left == curNode {
				return temp.Father
			}
			temp = temp.Father
		}
		return nil
	}
}
