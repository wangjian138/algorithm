package main

func main() {

}

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

func reverseLeftWords(s string, n int) string {
	var str []byte
	str = append(str, s[n-1:len(s)]...)
	str = append(str, s[0:n]...)
	return string(str)
}

// 切片是引用传递
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
