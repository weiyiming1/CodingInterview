package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	a = a[:3]
	fmt.Println(a)
}
