package main

func MaxInWindow(num []int, size int) []int {
	var maxInWin []int
	if len(num) >= size && size >= 1 {
		// 潜在的最大值索引
		var index []int
		for i := 0; i < size; i++ {
			for !isEmpty(&index) && num[i] >= num[getLast(&index)] {
				popBack(&index)
			}
			pushBack(&index, i)
		}
		for i := size; i < len(num); i++ {
			pushBack(&maxInWin, num[getFront(&index)])
			for !isEmpty(&index) && num[i] >= num[getLast(&index)] {
				popBack(&index)
			}
			if !isEmpty(&index) && getFront(&index) <= (i-size) {
				// 滑动窗口已经不包括index
				popFront(&index)
			}
			pushBack(&index, i)
		}
		pushBack(&maxInWin, num[getFront(&index)])
	}
	return maxInWin
}

func isEmpty(data *[]int) bool {
	// 是否为空
	return len(*data) == 0
}

func getLast(data *[]int) int {
	return (*data)[len(*data)-1]
}

func getFront(data *[]int) int {
	return (*data)[0]
}

func popBack(data *[]int) {
	*data = (*data)[:len(*data)-1]
}

func popFront(data *[]int) {
	*data = (*data)[1:]
}

func pushBack(data *[]int, new int) {
	*data = append(*data, new)
}
