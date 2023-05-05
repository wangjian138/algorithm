package main

func main() {

}

// https://leetcode.cn/problems/isomorphic-strings/?envType=study-plan-v2&id=top-interview-150
func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte, 0)

	s1Map := make(map[byte]byte, 0)

	for k, v := range s {
		s1, res := sMap[byte(v)]
		s2, res1 := s1Map[t[k]]
		if !res && !res1 {
			sMap[byte(v)] = t[k]
			s1Map[t[k]] = byte(v)
			continue
		}

		if res {
			if t[k] != s1 {
				return false
			}
		}

		if res1 {
			if s2 != byte(v) {
				return false
			}
		}

	}

	return true
}
