package main

import (
	"fmt"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func LinearSearch(array []int, t int) bool {
	i := 0
	for i < len(array) {
		if array[i] == t {
			return true
		}
		i++
	}
	return false
}

func BinarySearch(array []int, t int) bool {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] < t {
			left = mid + 1
		} else if array[mid] > t {
			right = mid - 1
		} else {
			return true
		}
	}

	return false
}

func RecursiveBinarySearch(array []int, t int, left int, right int) int {
	if left <= right {
		mid := (left + right) / 2
		if array[mid] < t {
			return RecursiveBinarySearch(array, t, mid+1, right)
		} else if array[mid] > t {
			return RecursiveBinarySearch(array, t, left, mid-1)
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	array := makeRange(0, 100000000)
	time1 := time.Now()
	bool := LinearSearch(array, 100000001)
	time2 := time.Now()
	fmt.Println("LinearSearch: ", time2.Sub(time1))
	fmt.Println("bool: ", bool)

	time3 := time.Now()
	bool1 := BinarySearch(array, 100000001)
	time4 := time.Now()
	fmt.Println("BinarySearch: ", time4.Sub(time3))
	fmt.Println("bool1: ", bool1)

	time5 := time.Now()
	bool2 := RecursiveBinarySearch(array, 100000001, 0, len(array)-1)
	time6 := time.Now()
	fmt.Println("RecursiveBinarySearch: ", time6.Sub(time5))
	fmt.Println("bool2: ", bool2)

}
