package main

import (
	"fmt"
	"testing"
)

func TestReverseSentence(t *testing.T) {
	chars := []rune{'i', ' ', 'a', 'm', ' ', 'a', ' ', 's', 't', 'u', 'd', 'e', 'n', 't'}
	demo := []rune{' ', 'a'}
	if demo[0] == ' ' {
		fmt.Println("空格")
	}
	ReverseSentence(&chars)
	fmt.Println(string(chars))
}
