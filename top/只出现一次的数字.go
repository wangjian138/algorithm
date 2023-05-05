package main

func main() {

}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * @param nums: List[int]
 * @return: int
 */
// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=136
