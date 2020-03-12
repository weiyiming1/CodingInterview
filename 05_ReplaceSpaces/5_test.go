package main

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	value := []rune{'i', ' ', 'a', 'm', ' ', 'o', 'k', ' ', '!', '\000'}
	orgin := make([]rune, 20)
	for i, j := range value {
		orgin[i] = j
	}
	new := ReplaceBlank(orgin, 20)
	fmt.Println(string(new))
}
