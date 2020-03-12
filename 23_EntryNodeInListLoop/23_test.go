package main

import (
	"fmt"
	"log"
	"testing"
)

func TestFindLoopEntry(t *testing.T) {
	f := Node{6, nil}
	e := Node{5, &f}
	d := Node{4, &e}
	c := Node{3, &d}
	b := Node{2, &c}
	a := Node{1, &b}
	f.Next = &c
	log.Println("输入链表")
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(d)
	log.Println(e)
	log.Println(f)
	entry := FindLoopEntry(&a)
	fmt.Println("环的入口节点：", entry)

}
