package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	pts := map[*ListNode]int{head: 0}

	pos := -1
	n := head.Next
	i := 1
	for n != nil {
		if i, ok := pts[n]; ok {
			pos = i
			fmt.Printf("i : %d\n", pos)
			fmt.Printf("n : %p\n", n)
			fmt.Println(pts)
			return true
		}
		pts[n] = i
		i++
		n = n.Next
	}
	fmt.Println(pts)
	return false
}
