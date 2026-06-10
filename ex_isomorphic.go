package main

//https://leetcode.com/problems/isomorphic-strings/description/?envType=study-plan-v2&envId=top-interview-150

func isIsomorphic(s string, t string) bool {
	m := map[rune]rune{}
	for i, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = rune(t[i])
		} else if m[r] != rune(t[i]) {
			return false
		}
	}
	return true
}
