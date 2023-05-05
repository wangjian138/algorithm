package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "baa"
	res := canConstruct(ransomNote, magazine)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/ransom-note/?envType=study-plan-v2&id=top-interview-150
func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := make(map[byte]int, 0)

	for _, v := range ransomNote {
		ransomNoteMap[byte(v)]++
	}

	for _, v := range magazine {
		num, res := ransomNoteMap[byte(v)]
		if !res {
			continue
		}
		num--
		ransomNoteMap[byte(v)] = num
		if num == 0 {
			delete(ransomNoteMap, byte(v))
		}
		if len(ransomNoteMap) == 0 {
			return true
		}
	}

	return false
}
