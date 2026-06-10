package main

import "strings"

//https://leetcode.com/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150

func wordPattern(pattern string, s string) bool {
	m := map[rune]string{}
	st := strings.Split(s, " ")
	for i, r := range pattern {
		if _, ok := m[r]; !ok {
			m[r] = st[i]
		} else if m[r] != st[i] {
			return false
		}
	}
	return true
}
