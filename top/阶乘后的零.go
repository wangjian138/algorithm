package main

func main() {

}

// https://leetcode.cn/problems/factorial-trailing-zeroes/?favorite=2ckc81c
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func trailingZeroes(n int) int {
	res := 0
	divisor := int64(5)
	for divisor <= int64(n) {
		res += n / int(divisor)
		divisor *= 5
	}
	return res
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=172
