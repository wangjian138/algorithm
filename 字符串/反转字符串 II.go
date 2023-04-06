package main

import "fmt"

func main() {
	s := "abcdefg"
	str := reverseStr(s, 8)
	fmt.Printf("str:%v\n", str)
}

// https://leetcode.cn/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	sByte := []byte(s)
	sByteLen := len(sByte)
	sByteNew := make([]byte, 0)

	if sByteLen <= k {
		return string(reverse([]byte(s)))
	}

	for i := 0; i < len(sByte); i += 2 * k {
		ik := i + k
		if ik > sByteLen {
			ik = sByteLen
		}
		reverse(sByte[i:ik])
	}

	return string(sByteNew)
}

func reverse(s []byte) []byte {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
