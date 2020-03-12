package main

import (
	"log"
	"testing"
)

func TestGetInversePairs(t *testing.T) {
	data := []int{7, 5, 6, 4}
	log.Println("逆序对:", GetInversePairs(&data))
}
