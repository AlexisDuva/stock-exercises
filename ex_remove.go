package main

func removeElementKeepOrder(nums []int, val int) int {
	k := len(nums)
	for i := 0; i < k; {
		if nums[i] == val {
			for j := i; j < k-1; j++ {
				nums[j] = nums[j+1]
			}
			k--
		} else {
			i++
		}
	}
	return k
}

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
