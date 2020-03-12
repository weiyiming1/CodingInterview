package main

import (
	"math/rand"
	"time"
)

// 解法1：
func generateRandom() int {
	var i int
	for i > 21 {
		i = 5*(rand5()-1) + rand5()
	}
	return i%7 + 1
}

/*
解法2：

func rand7() int {
    vals := [][]int{
       []int{1, 2, 3, 4, 5},
       []int{6, 7, 1, 2, 3},
       []int{4, 5, 6, 7, 1},
       []int{2, 3, 4, 5, 6},
       []int{7, 0, 0, 0, 0},
    }

    result := 0
    for result == 0 {
        i := rand5()
        j := rand5()
        result = vals[i-1][j-1]
    }
    return result
}*/

func rand5() int {
	seed := rand.NewSource(time.Now().Unix())
	r := rand.New(seed)
	return r.Intn(4) + 1
}
