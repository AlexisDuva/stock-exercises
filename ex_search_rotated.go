package main

import "fmt"

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min != max {
		mid := (max + min) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[max] == target {
			return max
		}
		if nums[min] < nums[mid] {
			if target >= nums[min] && target < nums[mid] {
				max = mid
			} else {
				min = mid
			}
		} else {
			if target > nums[mid] && target < nums[min] {
				min = mid
			} else {
				max = mid
			}
		}
		fmt.Printf("min : %d, mid : %d , max : %d\n", min, mid, max)
	}
	if min == target {
		return min
	} else {
		return -1
	}
}
