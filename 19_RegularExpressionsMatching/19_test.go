package main

import (
	"log"
	"testing"
)

func TestPatternMatch(t *testing.T) {
	str := []string{"a", "b", "c"}
	pattern := []string{"a", ".", "c"}
	flag := PatternMatch(str, pattern)
	log.Println(flag)
}
