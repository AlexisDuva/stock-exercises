package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	p1 := len(nums) - 1
	p2 := len(nums) - 2
	for p2 >= 0 {
		if nums[p2] >= p1-p2 {
			p1 = p2
		}
		p2--
		fmt.Printf("p1 : %d , p2 : %d\n", p1, p2)
	}
	return p1 == 0
}
