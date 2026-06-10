package main

import "slices"

//https://leetcode.com/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150

func minSubArrayLen(target int, nums []int) int {
	slices.Sort(nums)
	nb := removeDuplicates(nums)
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
