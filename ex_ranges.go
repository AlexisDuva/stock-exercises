package main

import "fmt"

//https://leetcode.com/problems/summary-ranges/?envType=study-plan-v2&envId=top-interview-150

func summaryRanges(nums []int) []string {
	r := []string{}
	p1 := 0
	p2 := 0
	for p1 < len(nums) {
		if p2 == len(nums) || nums[p2] != nums[p1]+(p2-p1) {
			if p1 == p2-1 {
				r = append(r, fmt.Sprintf("%d", nums[p1]))
			} else {
				r = append(r, fmt.Sprintf("%d->%d", nums[p1], nums[p2-1]))
			}
			p1 = p2
		}
		p2++
	}
	return r
}
