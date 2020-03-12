package main

import "fmt"

func main() {
	a := []string{"+213234"}
	fmt.Println(a)
	b := "-78891723"
	fmt.Println(b)
	fmt.Println(string(b[0]))
	fmt.Println(string(b[1]))
	fmt.Println(string(b[2]))
	/*for _, i := range b {
		fmt.Println(string(i))
	}*/

}
