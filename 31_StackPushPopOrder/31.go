package main

import "CodingInterview/getOffer-Go/lib"

func IsPopOrder(pushOrder, popOrder []int, len int) bool {
	flag := false
	if pushOrder != nil && popOrder != nil && len > 0 {
		nextPush := 0
		nextPop := 0

		// 先压入一些数据
		dataStack := lib.NewStack()
		for nextPop < len {
			// 当辅助栈的栈顶元素不是要弹出的元素
			for dataStack.Empty() || dataStack.Top() != popOrder[nextPop] {
				// 如果所有数字都压入辅助栈了，退出循环
				if nextPush == len {
					break
				}
				//压入数字入栈
				dataStack.Push(pushOrder[nextPush])
				nextPush++
			}
			// 没有匹配成功
			if dataStack.Top() != popOrder[nextPop] {
				break
			}
			dataStack.Pop()
			nextPop++
		}
		if dataStack.Empty() && nextPop == len {
			flag = true
		}
	}
	return flag
}
