package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

func printPrime() {
	i := 1
	for i > 0 && i < 50 {
		if isPrime(i) {
			fmt.Println(i)
		}
		i++
	}
}

func main() {
	printPrime()
}
