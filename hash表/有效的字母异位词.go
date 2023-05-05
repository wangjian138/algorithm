package main

func main() {

}

// https://leetcode.cn/problems/valid-anagram/?envType=study-plan-v2&id=top-interview-150
func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int, 0)
	for _, v := range s {
		sMap[byte(v)]++
	}

	for _, v := range t {
		s1, res := sMap[byte(v)]
		if !res {
			return false
		}
		s1--
		sMap[byte(v)] = s1
		if s1 == 0 {
			delete(sMap, byte(v))
		}
	}
	if len(sMap) == 0 {
		return true
	}

	return false
}
