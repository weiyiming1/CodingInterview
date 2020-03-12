package main

import (
	"log"
	"testing"
)

func TestGetNextNode(t *testing.T) {
	h := &BinaryTreeNode{"h", nil, nil, nil}
	i := &BinaryTreeNode{"i", nil, nil, nil}
	d := &BinaryTreeNode{"d", nil, nil, nil}
	e := &BinaryTreeNode{"e", h, i, nil}
	f := &BinaryTreeNode{"f", nil, nil, nil}
	g := &BinaryTreeNode{"g", nil, nil, nil}
	b := &BinaryTreeNode{"b", d, e, nil}
	c := &BinaryTreeNode{"c", f, g, nil}
	a := &BinaryTreeNode{"a", b, c, nil}

	// 添加父节点
	b.Father = a
	c.Father = a
	d.Father = b
	e.Father = b
	f.Father = c
	g.Father = c
	h.Father = e
	i.Father = e

	log.Println("测试节点：", b.Value)
	log.Println("中序遍历的下一个节点：", GetNextNode(b).Value)

}
