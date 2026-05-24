package main

// func groupAnagrams(strs []string) [][]string {
// 	res := [][]string{}
// 	for i := 0; i < len(strs); {
// 		s1 := strs[i]
// 		fmt.Printf("s1 :%s\n", s1)
// 		res = append(res, []string{})
// 		sum1 := 0
// 		m1 := map[rune]int{}
// 		for _, r := range s1 {
// 			m1[r] = m1[r] + 1
// 			sum1++
// 		}
// 		for j := i; j < len(strs); j++ {
// 			s2 := strs[j]
// 			fmt.Printf("s2 :%s\n", s2)
// 			isAnagram := true
// 			sum2 := sum1
// 			m2 := maps.Clone(m1)
// 			for _, r := range s2 {
// 				m2[r] = m2[r] - 1
// 				sum2--
// 				for k, v := range m2 {
// 					fmt.Printf("%c : %d\n", k, v)
// 				}
// 				fmt.Printf("sum2 :%d\n", sum2)
// 				//Check in s2 if too much of one letter OR a lettre not present in s1
// 				if m2[r] < 0 {
// 					isAnagram = false
// 					break
// 				}
// 			}
// 			//Check if not enough of correct letters in s2
// 			if sum2 > 0 {
// 				isAnagram = false
// 			}
// 			if isAnagram {
// 				fmt.Printf("%s is anagram of  %s \n", s2, s1)
// 				res[i] = append(res[i], s2)
// 			}
// 		}
// 	}
// 	return res
// }

// func groupAnagrams(strs []string) [][]string {
// 	res := [][]string{}
// 	ms := []map[rune]int{}
// 	for i, s := range strs {
// 		for _, r := range s {
// 			ms[i][r] = ms[i][r] + 1
// 		}

// 	}
// 	p1 := 0
// 	p2 := 1
// 	for p1 < len(strs) {
// 		for p2 := p1 + 1; p2 < len(strs); {
// 			if maps.Equal(ms[p1], ms[p2]) {
// 				res[p1] = append(res[p1], strs[p2])
// 				slices.Delete(strs, p2, p2+1)
// 			}
// 		}
// 	}
// 	return res
// }
