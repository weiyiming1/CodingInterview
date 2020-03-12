package main

import (
	"CodingInterview/getOffer-Go/lib"
	"log"
	"testing"
)

func TestSymmetrical(t *testing.T) {
	d := lib.TreeNode{5, nil, nil}
	e := lib.TreeNode{7, nil, nil}
	f := lib.TreeNode{7, nil, nil}
	g := lib.TreeNode{5, nil, nil}
	b := lib.TreeNode{6, &d, &e}
	c := lib.TreeNode{6, &f, &g}
	a := lib.TreeNode{8, &b, &c}
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(d)
	log.Println(e)
	log.Println(f)
	log.Println(g)
	flag := Symmetrical(&a)
	log.Println(flag)

	dd := lib.TreeNode{5, nil, nil}
	ee := lib.TreeNode{7, nil, nil}
	ff := lib.TreeNode{7, nil, nil}
	gg := lib.TreeNode{5, nil, nil}
	bb := lib.TreeNode{6, &dd, &ee}
	cc := lib.TreeNode{9, &ff, &gg}
	aa := lib.TreeNode{8, &bb, &cc}
	log.Println(aa)
	log.Println(bb)
	log.Println(cc)
	log.Println(dd)
	log.Println(ee)
	log.Println(ff)
	log.Println(gg)
	flag = Symmetrical(&aa)
	log.Println(flag)

	ddd := lib.TreeNode{7, nil, nil}
	eee := lib.TreeNode{7, nil, nil}
	fff := lib.TreeNode{7, nil, nil}
	bbb := lib.TreeNode{7, &ddd, &eee}
	ccc := lib.TreeNode{7, &fff, nil}
	aaa := lib.TreeNode{7, &bbb, &ccc}
	log.Println(aaa)
	log.Println(bbb)
	log.Println(ccc)
	log.Println(ddd)
	log.Println(eee)
	log.Println(fff)
	flag = Symmetrical(&aaa)
	log.Println(flag)
}
