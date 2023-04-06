package main

import "fmt"

func main() {
	s := "baa"
	p := "aa"
	allIndex := findAnagrams(s, p)
	fmt.Printf("allIndex:%v\n", allIndex)
}

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
func findAnagrams(s string, p string) []int {
	pMap := make(map[byte]int, len(p))
	for v := range p {
		pMap[p[v]] = pMap[p[v]] + 1
	}

	allIndex := make([]int, 0)

	windowMap := make(map[byte]int, len(s))

	pMapLen := len(pMap)
	pLen := len(p)

	sLen := len(s)

	left := 0
	right := 0
	valid := 0

	for right < sLen {
		rByte := s[right]
		windowMap[rByte] = windowMap[rByte] + 1
		right++

		val, res := pMap[rByte]
		if res {
			if val == windowMap[rByte] {
				valid++
			}
		}

		for valid == pMapLen {
			lByte := s[left]

			val, res := pMap[lByte]
			if res {
				if val == windowMap[lByte] {
					if right-left == pLen {
						allIndex = append(allIndex, left)
					}
					valid--
				}
			}
			windowMap[lByte] = windowMap[lByte] - 1
			left++
		}
	}

	return allIndex
}
