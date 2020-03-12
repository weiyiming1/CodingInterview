package main

import (
	"log"
	"testing"
)

func TestGetLastCommonParent(t *testing.T) {
	F := Node{6, nil}
	G := Node{7, nil}
	H := Node{8, nil}
	I := Node{9, nil}
	J := Node{10, nil}
	D := Node{4, []*Node{&F, &G}}
	E := Node{5, []*Node{&H, &I, &J}}
	B := Node{2, []*Node{&D, &E}}
	C := Node{3, nil}
	A := Node{1, []*Node{&B, &C}}
	pLast := GetLastCommonParent(&A, &F, &H)
	if pLast != nil {
		log.Printf("最低公共节点:%d, %v", pLast.val, pLast.children)
	} else {
		log.Println("没有最低公共节点")
	}
}
