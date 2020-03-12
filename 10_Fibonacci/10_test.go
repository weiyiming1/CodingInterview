package main

import (
	"log"
	"testing"
)

func TestFibonacci(t *testing.T) {
	sum := Fibonacci(3)
	log.Println(sum)
	sum = Fibonacci(4)
	log.Println(sum)
	sum = Fibonacci(5)
	log.Println(sum)
	sum = Fibonacci(6)
	log.Println(sum)
	sum = Fibonacci(7)
	log.Println(sum)
}
