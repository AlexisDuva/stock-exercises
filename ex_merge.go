package main

//Import Stock from ex1_1.go
import (
	"fmt"
	"slices"
)

// func Shift(nums *[]int, start int, end int) {
// 	for i := end - 1; i >= start; i-- {
// 		(*nums)[i+1] = i
// 	}
// }

// func InsertSorted(nums *[]int, start int, end int, v int, size int) {
// 	mid := (end + start) / 2
// 	fmt.Printf("start:%d  mid:%d  end:%d \n", start, mid, end)
// 	if v < (*nums)[mid] {
// 		if (end - start) == 1 {
// 			Shift(nums, mid, size)
// 			(*nums)[mid] = v
// 			fmt.Printf("nums: %d \n", *nums)
// 		} else {
// 			InsertSorted(nums, start, mid, v, size)
// 		}
// 	} else if v > (*nums)[mid] {
// 		if (end - start) == 1 {
// 			Shift(nums, mid+1, size)
// 			(*nums)[mid+1] = v
// 			fmt.Printf("nums: %d \n", *nums)
// 		} else {
// 			InsertSorted(nums, mid, end, v, size)
// 		}
// 	} else {
// 		Shift(nums, mid+1, size)
// 		(*nums)[mid+1] = v
// 		fmt.Printf("nums: %d \n", *nums)
// 	}
// }

// func merge(nums1 *[]int, m int, nums2 []int, n int) {
// 	for i, v := range nums2 {
// 		fmt.Printf("merge() :v: %d \n", v)
// 		InsertSorted(nums1, 0, m+i, v)
// 		fmt.Printf("merge() : nums1: %d \n", *nums1)
// 	}
// }

func insertSimple(vals []int) {
	vals2 := slices.Insert(vals, 1, 7)
	fmt.Printf("vals2 : %d \n", vals2)
	fmt.Printf("len(vals2) : %d \n", len(vals2))
	fmt.Printf("cap(vals2) : %d \n", cap(vals2))
	// for i, v := range vals2 {
	// 	vals[i] = v
	// }
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1

	for i := m + n - 1; i >= 0; i-- {
		// Check if we've run out of elements in nums2
		if p2 < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}