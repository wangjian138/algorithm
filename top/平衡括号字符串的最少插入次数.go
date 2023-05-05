package main

func main() {

}

// https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string/
func minInsertions(s string) int {
	res := 0
	need := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			need += 2
			if need%2 == 1 {
				res++
				need--
			}
		} else if s[i] == ')' {
			need--
			if need == -1 {
				res++
				need = 1
			}
		}
	}

	return res + need
}

// https://labuladong.github.io/article/?qno=1541
