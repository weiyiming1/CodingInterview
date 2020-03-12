package main

// 普通树（多个子树），子树没有指向父节点的指针
type Node struct {
	val      int
	children []*Node
}

func GetLastCommonParent(pRoot, pNode1, pNode2 *Node) *Node {
	if pRoot == nil || pNode1 == nil || pNode2 == nil {
		return nil
	}
	path1 := make([]*Node, 1)
	path2 := make([]*Node, 1)
	getNodePath(pRoot, pNode1, &path1)
	getNodePath(pRoot, pNode2, &path2)
	return getLastCommonNode(&path1, &path2)
}

func getNodePath(pRoot, pNode *Node, path *[]*Node) bool {
	if pRoot == pNode {
		return true
	}
	*path = append(*path, pRoot)
	isFound := false
	var idx int
	for !isFound && idx < len(pRoot.children) {
		isFound = getNodePath(pRoot.children[idx], pNode, path)
		idx++
	}
	if !isFound { // 没有找到，弹出当前节点
		*path = (*path)[:len(*path)-1]
	}
	return isFound
}

func getLastCommonNode(path1, path2 *[]*Node) *Node {
	var pLast *Node
	for i := 0; i < len(*path1); i++ {
		if (*path1)[i] != (*path2)[i] {
			break
		}
		pLast = (*path1)[i]
	}
	return pLast
}
