package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestGetMirror(t *testing.T) {
	d := lib.TreeNode{5, nil, nil}
	e := lib.TreeNode{7, nil, nil}
	f := lib.TreeNode{9, nil, nil}
	g := lib.TreeNode{11, nil, nil}
	b := lib.TreeNode{6, &d, &e}
	c := lib.TreeNode{10, &f, &g}
	a := lib.TreeNode{8, &b, &c}

	log.Println("输入：")
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(d)
	log.Println(e)
	log.Println(f)
	log.Println(g)
	GetMirror(&a)
	log.Println("输出：")
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(d)
	log.Println(e)
	log.Println(f)
	log.Println(g)
}
