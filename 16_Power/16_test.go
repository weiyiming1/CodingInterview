package main

import (
	"log"
	"testing"
)

func TestPowerWithExponent(t *testing.T) {
	res := PowerWithExponent(3, 3)
	log.Println(res)
	res = PowerWithExponent(2, 6)
	log.Println(res)
	res = PowerWithExponent(6, 3)
	log.Println(res)
}
