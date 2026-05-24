package main

import (
	"fmt"
	"math"
	"strconv"
)

func isHappy(n int) bool {
	m := map[int]int{}
	sum := n
	for i := 0; ; i++ {
		sum2 := 0
		for _, r := range strconv.Itoa(sum) {
			sum2 += int(math.Pow(float64(r-'0'), 2))
		}
		fmt.Println(sum2)
		if sum2 == 1 {
			return true
		}
		if _, ok := m[sum2]; ok {
			return false
		}
		m[sum2] = i
		sum = sum2
	}
}
