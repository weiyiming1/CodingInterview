package main

func FindGreatestSumOfSubArray(data []int) *int {
	if len(data) == 0 {
		return nil
	}

	nCurSum := 0
	nGreatestSum := 0
	for _, val := range data {
		if nCurSum <= 0 {
			nCurSum = val
		} else {
			nCurSum += val
		}

		if nCurSum > nGreatestSum {
			nGreatestSum = nCurSum
		}
	}
	return &nGreatestSum
}
