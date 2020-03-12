package main

import (
	"fmt"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	chars := []rune{'a', 'r', 'a', 'b', 'c', 'a', 'c', 'f', 'r'}
	len := LongestSubString(chars)
	fmt.Println(len)

	chars = []rune{'r', 'f', 'r'}
	len = LongestSubString(chars)
	fmt.Println(len)
}
