package main

import "slices"

func fusionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		return interSort(fusionSort(nums[:len(nums)/2]), fusionSort(nums[len(nums)/2:]))
	}
}

func interSort(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	if nums1[0] <= nums2[0] {
		return slices.Concat(nums1[0:1], interSort(nums1[1:], nums2))
	} else {
		return slices.Concat(nums2[0:1], interSort(nums1, nums2[1:]))
	}
}
