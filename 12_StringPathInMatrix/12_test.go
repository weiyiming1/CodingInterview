package main

import (
	"log"
	"testing"
)

func TestHasPath(t *testing.T) {
	matrix := [][]rune{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}}
	rows := len(matrix)
	cols := len(matrix)
	char1 := []rune{'b', 'f', 'c', 'e'}
	res1 := HasPath(matrix, rows, cols, char1)
	log.Println(res1)

	char2 := []rune{'a', 'b', 'f', 'b'}
	res2 := HasPath(matrix, rows, cols, char2)
	log.Println(res2)
}
