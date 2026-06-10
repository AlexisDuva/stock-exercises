package main

import "fmt"

//https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-interview-150

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	p1 := 0
	p2 := 1
	k := 1
	for p2 < len(s) && len(s)-p1 > k {
		for i := p1; i < p2; i++ {
			if s[i] == s[p2] {
				p1 = i + 1
				break
			}
		}
		k = max(k, p2-p1+1)
		p2++
	}
	fmt.Printf("LongestSubString : %s\n", s[p1:p2])
	return k
}
