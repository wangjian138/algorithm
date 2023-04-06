package main

import (
	"fmt"
)

func main() {
	wStr := minWindow("abc", "ac")
	fmt.Printf("wStr:%v\n", wStr)
}

// https://leetcode.cn/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	tMap := make(map[byte]int, len(t))
	for k := range t {
		tMap[t[k]] = tMap[t[k]] + 1
	}
	tMapLen := len(tMap)
	tLen := len(t)

	windowMap := make(map[byte]int, len(s))

	// 最小的长度
	minLen := len(s)
	start := -1

	right := 0
	left := 0
	size := len(s)
	valid := 0

	for right < size {
		// 判断
		rightByte := s[right]
		right++

		windowMap[rightByte] = windowMap[rightByte] + 1
		val, res := tMap[rightByte]
		if res {
			if val == windowMap[rightByte] {
				valid++
			}
		}

		// 说明此时满足进行收缩
		for valid == tMapLen {
			// 出现了最小字符串
			if right-left == tLen {
				return s[left:right]
			}
			if right-left <= minLen {
				minLen = right - left
				start = left

			}
			leftByte := s[left]
			val, res := tMap[leftByte]
			if res {
				if val == windowMap[leftByte] {
					valid--
				}
				windowMap[leftByte] = windowMap[leftByte] - 1
			}
			left++
		}
	}

	if start >= 0 {
		return s[start : start+minLen]
	}

	return ""
}

func minWindow1(s string, t string) string {

	if s == "" || t == "" {
		return ""
	}

	if s == t {
		return s
	}

	sByte := []byte(s)
	sByteLen := len(sByte)

	tByte := []byte(t)

	if len(tByte) > sByteLen {
		return ""
	}

	tByteMap := make(map[byte]int, 0)

	for _, tb := range tByte {
		tByteMap[tb] = tByteMap[tb] + 1
	}

	tByteLen := len(tByteMap)

	windowMap := make(map[byte]int, 0)

	right := 0
	left := -1
	valid := 0
	start := -1

	valLen := len(sByte)

	for right <= len(sByte)-1 {
		val := sByte[right]
		right++

		_, res := tByteMap[val]
		if res == true {
			windowMap[val] = windowMap[val] + 1

			if left == -1 {
				left = right - 1
			}

			if windowMap[val] == tByteMap[val] {
				valid++
			}
		}

		for valid == tByteLen {
			if right-left <= valLen {
				start = left
				valLen = right - left
			}

			val := sByte[left]

			left++

			_, res := tByteMap[val]
			if res == true {
				windowVal, _ := windowMap[val]
				if windowMap[val] == tByteMap[val] {
					valid--
				}
				windowMap[val] = windowVal - 1
			}
			if valLen == len(tByte) {
				goto end
			}
		}

	}

end:
	str := ""
	if start >= 0 {
		strings := sByte[start : start+valLen]
		str = string(strings)
	}

	return str
}
