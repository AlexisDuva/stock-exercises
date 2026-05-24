package main

import "strings"

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
