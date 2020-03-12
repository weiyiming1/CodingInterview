package main

func MaxCut(len int) int {
	//因为要求长度n>1,所以这里返回0表示输入非法
	if len < 2 {
		return 0
	}
	//长度为2时,因为要求剪下段数m>1,所以最大是1x1=1
	if len == 2 {
		return 1
	}
	//长度为3时,因为要求剪下段数m>1,所以最大是1x2=2
	if len == 3 {
		return 2
	}
	//运行至此,说明绳子的长度是>3的,这之后0/1/2/3这种子问题最大就是其自身长度
	//而不再需要考虑剪一刀的问题,因为它们剪一刀没有不剪来的收益高
	//而在当下这么长的绳子上剪过才可能生成0/1/2/3这种长度的绳子,它们不需要再减
	//所以下面的表中可以看到它们作为子问题的值和上面实际返回的是不一样的

	//在表中先打上子绳子的最大乘积
	res := make([]int, len+1)
	res[0] = 0
	res[1] = 1
	res[2] = 2
	res[3] = 3 //到3为止都是不剪最好

	for i := 4; i <= len; i++ {
		max := 0
		for j := 0; j <= i/2; j++ {
			temp := res[j] * res[i-j]
			if max < temp {
				max = temp
			}
		}
		res[i] = max
	}
	return res[len]
}
