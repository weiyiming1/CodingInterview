package main

func SwitchMin(nums []int) int {
	if &nums == nil || len(nums) == 0 {
		return -1
	}

	idx1 := 0
	idx2 := len(nums) - 1
	idxMid := idx1
	for nums[idx1] >= nums[idx2] {
		// 正常情况下，终止条件
		if idx2-idx1 == 1 {
			idxMid = idx2
			break
		}
		idxMid = (idx1 + idx2) / 2
		if nums[idx1] == nums[idx2] && nums[idxMid] == nums[idx1] {
			return minInOrder(nums, idx1, idx2)
		}
		if nums[idxMid] >= nums[idx1] {
			idx1 = idxMid
		} else if nums[idxMid] <= nums[idx2] {
			idx2 = idxMid
		}
	}
	return nums[idxMid]
}

func minInOrder(nums []int, idx1, idx2 int) int {
	res := nums[idx1]
	for i := idx1 + 1; i <= idx2; i++ {
		if res > nums[i] {
			res = nums[i]
		}
	}
	return res
}
