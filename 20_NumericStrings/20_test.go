package main

import (
	"log"
	"testing"
)

func TestIsNumberic(t *testing.T) {
	a := "3e+4"
	flag := IsNumberic(a)
	log.Println(flag)

	a = "3e+4.4"
	flag = IsNumberic(a)
	log.Println(flag)

	a = "-3e4.4"
	flag = IsNumberic(a)
	log.Println(flag)

	a = "-3e5"
	flag = IsNumberic(a)
	log.Println(flag)

	a = "-3"
	flag = IsNumberic(a)
	log.Println(flag)

	a = "1.2.3"
	flag = IsNumberic(a)
	log.Println(flag)
}
