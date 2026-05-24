package main

import "slices"

func minSubArrayLen(target int, nums []int) int {
	slices.Sort(nums)
	nb := RemoveDuplicates(nums)
	sum := 0
	p := nb - 1
	for sum < target && p >= 0 {
		sum += nums[p]
		p--
	}
	if sum < target {
		return 0
	}
	return nb - p - 1
}
