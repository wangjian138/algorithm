package main

import "fmt"

func main() {
	valid := isValid("{}[{[}]")
	fmt.Printf("valid:%v\n", valid)
}

// https://leetcode.cn/problems/valid-parentheses/
func isValid(s string) bool {
	validArr := make([]byte, 0)
	for _, sub := range s {
		if sub == '{' || sub == '[' || sub == '(' {
			validArr = append(validArr, byte(sub))
			continue
		}
		if len(validArr) == 0 {
			return false
		}
		lastStr := validArr[len(validArr)-1]
		validArr = validArr[0 : len(validArr)-1]
		if sub == '}' {
			if lastStr != '{' {
				return false
			}
		} else if sub == ']' {
			if lastStr != '[' {
				return false
			}
		} else if sub == ')' {
			if lastStr != '(' {
				return false
			}
		}
	}
	if len(validArr) > 0 {
		return false
	}
	return true
}
