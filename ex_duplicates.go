package main

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicates(nums []int) int {
	k := 0
	p1 := 0
	p2 := 1
	for p2 < len(nums) {
		if nums[p1] != nums[p2] {
			nums[k] = nums[p1]
			p1 = p2
			k++
		}
		p2++
	}
	nums[k] = nums[p1]
	k++
	return k
}
