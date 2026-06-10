package main

//https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150

func canConstruct(ransomNote string, magazine string) bool {
	lettersCount := map[rune]int{}
	for _, c := range magazine {
		lettersCount[c] = lettersCount[c] + 1
	}
	for _, c := range ransomNote {
		lettersCount[c] = lettersCount[c] - 1
		if lettersCount[c] < 0 {
			return false
		}
	}
	return true
}
