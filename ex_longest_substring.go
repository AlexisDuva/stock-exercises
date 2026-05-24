package main

import "fmt"

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

func lengthOfLongestSubstringWithHashMap(s string) int {
	if len(s) < 1 {
		return 0
	}
	seen := make(map[byte]int)
	p1 := 0
	p2 := 1
	k := 1
	for ; p2 < len(s) && len(s)-p1 > k; p2++ {
		if idx, ok := seen[s[p2]]; ok && idx >= p1 {
			p1 = idx + 1
		}
		seen[s[p2]] = p2
		k = max(k, p2-p1+1)
	}
	fmt.Printf("LongestSubString : %s\n", s[p1:p2])
	return k
}
