package main

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			if v1+v2 == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}

// func twoSumMap(nums []int, target int) []int {
// 	sum := map[int]int{}
// 	sum[nums[0]] = target
// 	for i, v := range nums[1:] {
// 		if map[]
// 		sum[v] = target
// 	}
// 	return []int{}
// }
