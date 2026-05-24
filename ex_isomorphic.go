package main

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
