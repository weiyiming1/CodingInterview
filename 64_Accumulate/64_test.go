package main

import (
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	res := Sum(4)
	log.Println("res:", res)
}
