package main

import (
	"log"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance := GetInstance()
	if instance != nil {
		log.Println("generated using sync.once")
	} else {
		log.Println("called before")
	}

	instance2 := GetInstance()
	if instance2 != instance {
		log.Println("generated using sync.once")
	} else {
		log.Println("called before")
	}
}
