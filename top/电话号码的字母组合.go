package main

import "fmt"

func main() {
	digits := "23"
	res := letterCombinations(digits)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	res = make([]string, 0)
	return getCombination(digits)
}

var res []string

var digits2Letter = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "y", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func getCombination(digits string) []string {
	if digits == "" {
		return res
	}

	strArr := digits2Letter[digits[0]]
	var newRes []string
	if len(res) == 0 {
		res = strArr
	} else {
		for _, str := range strArr {
			for _, v := range res {
				newRes = append(newRes, v+str)
			}
		}
		res = newRes
	}

	getCombination(digits[1:])
	return res
}
