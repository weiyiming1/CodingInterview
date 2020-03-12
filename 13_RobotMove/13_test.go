package main

import (
	"log"
	"testing"
)

func TestMovingCount(t *testing.T) {
	log.Println(MovingCount(5, 10, 10))  // 21
	log.Println(MovingCount(15, 20, 20)) // 359
}
