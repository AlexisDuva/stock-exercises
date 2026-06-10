package main

//https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150

func removeElement(nums []int, val int) int {
	p1 := 0
	p2 := len(nums) - 1
	for p1 < p2 {
		if nums[p1] == val {
			for nums[p2] == val && p2 > p1 {
				p2--
			}
			nums[p1] = nums[p2]
			p2--
		}
		p1++
	}
	return p2 + 1
}
