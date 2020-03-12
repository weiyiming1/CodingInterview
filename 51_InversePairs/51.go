package main

func GetInversePairs(pData *[]int) int {
	if pData == nil || len(*pData) <= 0 {
		return 0
	}
	copy := make([]int, len(*pData))
	count := InversePairs(pData, &copy, 0, len(*pData)-1)
	return count
}

func InversePairs(pData, temp *[]int, start, end int) int {
	if start == end {
		(*temp)[start] = (*pData)[start]
		return 0
	}
	subLen := (end - start) / 2

	// ⚠️：这里的pData, temp位置不一样
	left := InversePairs(temp, pData, start, start+subLen)
	right := InversePairs(temp, pData, start+subLen+1, end)

	i := start + subLen //前半段最后一个
	j := end            //后半段最后一个
	idxCopy := end      //开始拷贝的位置
	count := 0          // 逆序数
	for i >= start && j >= start+subLen+1 {
		if (*pData)[i] > (*pData)[j] {
			(*temp)[idxCopy] = (*pData)[i]
			idxCopy--
			i--
			count += j - start - subLen // 对应的逆序对
		} else {
			(*temp)[idxCopy] = (*pData)[j]
			idxCopy--
			j--
		}
	}
	for i >= start {
		(*temp)[idxCopy] = (*pData)[i]
		idxCopy--
		i--
	}
	for j >= start+subLen+1 {
		(*temp)[idxCopy] = (*pData)[j]
		idxCopy--
		j--
	}
	return left + right + count
}
