package main

import "fmt"

func main() {
	s := "  the sky is blue"
	res := reverseWords(s)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	fmt.Printf("reverseWords b:%v\n", string(b))

	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	fmt.Printf("reverseWords b:%v\n", string(b))

	//2.反转整个字符串
	reverse(b)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(b[i:j])
		i = j
		i++
	}
	return string(b)
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
