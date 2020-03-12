package main

import (
	"log"
	"testing"
)

func TestGetTranslation(t *testing.T) {
	log.Println(GetTranslation(-10))
	log.Println(GetTranslation(1234))
	log.Println(GetTranslation(12258))
}
