面试题21：调整数组顺序使奇数位于偶数前面
> 题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有
奇数位于数组的前半部分，所有偶数位于数组的后半部分。

```
package main

import (
    "fmt"
)

func main() {
    callback(1, Add)
}

func Add(a, b int) {
    fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
    f(y, 2) // this becomes Add(1, 2)
}
```
输出：

`The sum of 1 and 2 is: 3`