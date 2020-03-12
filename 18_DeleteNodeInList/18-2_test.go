package main

import (
	"log"
	"testing"
)

func TestDeleteDuplicate(t *testing.T) {
	c := Node{
		2,
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

	head := Node{
		-1,
		&a,
	}
	log.Println("输入节点：")
	log.Println("\t", *head.Next)
	log.Println("\t", *(head.Next.Next))
	log.Println("\t", *(head.Next.Next.Next))
	DeleteDuplicate(&head)
	log.Println("处理后节点：")
	log.Println("\t", *head.Next)
	if head.Next.Next != nil {
		t.Errorf("got %v, expected nil", head.Next.Next)
	}
	d := Node{
		4,
		nil,
	}
	head = Node{
		-1,
		&d,
	}
	log.Println("")
	log.Println("输入节点：")
	log.Println("\t", *head.Next)
	DeleteDuplicate(&head)
	log.Println("处理后节点：")
	log.Println("\t", *head.Next)

	f := Node{
		5,
		nil,
	}
	e := Node{
		5,
		&f,
	}
	head = Node{
		-1,
		&e,
	}
	log.Println("")
	log.Println("输入节点：")
	log.Println("\t", *head.Next)
	log.Println("\t", *head.Next.Next)
	DeleteDuplicate(&head)
	log.Println("处理后节点：")
	if head.Next != nil {
		t.Errorf("got %v, expected nil", head.Next)
	}

}
