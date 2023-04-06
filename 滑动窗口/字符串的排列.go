package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidboaoo"
	isInclude := checkInclusion(s1, s2)
	fmt.Printf("isInclude:%v\n", isInclude)
}

// https://leetcode.cn/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	s1, s2 = s2, s1

	s2Map := make(map[byte]int, 0)

	for k := range s2 {
		s2Map[s2[k]] = s2Map[s2[k]] + 1
	}

	windowMap := make(map[byte]int)
	s2MapLen := len(s2Map)
	s2Len := len(s2)

	s1Len := len(s1)

	valid := 0

	left := 0
	right := 0
	for right < s1Len {
		rByte := s1[right]
		right++

		windowMap[rByte] = windowMap[rByte] + 1

		val, res := s2Map[rByte]
		if res {
			if val == windowMap[rByte] {
				valid++
			}
		}
		for valid == s2MapLen {
			lByte := s1[left]
			val, res := s2Map[lByte]
			if res {
				if val == windowMap[lByte] {
					if right-left == s2Len {
						return true
					}
					valid--
				}
				windowMap[lByte] = windowMap[lByte] - 1
			}
			left++
		}
	}

	return false
}
