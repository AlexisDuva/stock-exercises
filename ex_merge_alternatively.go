package main

import (
	"strings"
)

//https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	i := 0
	for i = range min(len(word1), len(word2)) {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[i])
	}
	if len(word1) < len(word2) {
		sb.WriteString(word2[i+1:])
	} else if len(word1) > len(word2) {
		sb.WriteString(word1[i+1:])
	}
	return sb.String()
}
