package main

import (
	"log"
	"testing"
)

func TestFirstNotRepeatChar(t *testing.T) {
	s := "abaccdeff"
	demo := FirstNotRepeatChar(s)
	log.Println(demo)
}
