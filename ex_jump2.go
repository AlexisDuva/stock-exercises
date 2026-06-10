package main

import "fmt"

//https://leetcode.com/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150

func jump(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	p2 := len(nums) - 2
	steps := []int{len(nums) - 1}
	for p2 >= 0 {
		for i, s := range steps {
			if p2+nums[p2] >= s {
				steps = append(steps[:i+1], p2)
				break
			}
		}
		p2--
		fmt.Printf("steps : %d , p2 : %d\n", steps, p2)
	}
	return len(steps) - 1
}
