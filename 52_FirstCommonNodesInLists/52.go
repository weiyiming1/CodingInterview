package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func FindFirstCommonNode(pNode1, pNode2 *Node) *Node {

	var len1, len2, dis int
	pN1 := pNode1
	pN2 := pNode2
	// 获取两个链的长度
	for true {
		len1++
		if pN1.next != nil {
			pN1 = pN1.next
		} else {
			break
		}
	}

	for true {
		len2++
		if pN2.next != nil {
			pN2 = pN2.next
		} else {
			break
		}
	}
	fmt.Println("len1", len1)
	fmt.Println("len2", len2)

	var pLong, pShort *Node
	if len1 > len2 {
		dis = len1 - len2
		pLong = pNode1
		pShort = pNode2
	} else {
		dis = len2 - len1
		pLong = pNode2
		pShort = pNode1
	}

	// 在长链上先走几步
	for i := 0; i < dis; i++ {
		pLong = pLong.next
	}

	//再同时遍历
	for pLong != nil && pShort != nil && pLong != pShort {
		pLong = pLong.next
		pShort = pShort.next
	}

	// 得到的第一个公共节点
	return pLong
}
