package main

import (
	"log"
	"testing"
)

func TestClone(t *testing.T) {
	e := ComplexListNode{5, nil, nil}
	d := ComplexListNode{4, &e, nil}
	c := ComplexListNode{3, &d, nil}
	b := ComplexListNode{2, &c, nil}
	a := ComplexListNode{1, &b, nil}

	a.Sibling = &c
	b.Sibling = &e
	d.Sibling = &b

	log.Println(a)
	log.Println(a.Next)
	log.Println(a.Next.Next)
	log.Println(a.Next.Next.Next)
	log.Println(a.Next.Next.Next.Next)

	aa := Clone(&a)
	log.Println(aa)
	log.Println(aa.Next)
	log.Println(aa.Next.Next)
	log.Println(aa.Next.Next.Next)
	log.Println(aa.Next.Next.Next.Next)
}
