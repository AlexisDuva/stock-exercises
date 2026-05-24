package main

import "fmt"

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min != max {
		mid := (max + min) / 2
		if target < nums[mid] {
			max = mid
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			return mid // target found
		}
		fmt.Printf("min : %d, mid : %d , max : %d\n", min, mid, max)
	}
	//target not found
	if target > nums[min] {
		return min + 1
	} else {
		return min
	}
}
