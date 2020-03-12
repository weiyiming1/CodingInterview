package main

import (
	"log"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	c := Node{
		3,
		nil,
	}
	b := Node{
		2,
		&c,
	}
	a := Node{
		1,
		&b,
	}

	log.Println(a)
	log.Println(*a.Next)
	log.Println(*a.Next.Next)

	log.Println("------")
	DeleteNode(&a, &b)
	log.Println("删除中间节点")
	log.Println(a)
	log.Println(*a.Next)
	log.Println("")
	d := Node{
		111,
		nil,
	}
	log.Println("当前链表：", d)
	DeleteNode(&d, &d)
	log.Println("删除首节点")
	log.Println(d)
}
