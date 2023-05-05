package main

func main() {

}

// https://leetcode.cn/problems/gas-station/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// 相当于图像中的坐标点和最低点
	sum, minSum := 0, 0
	start := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			// 经过第 i 个站点后，使 sum 到达新低
			// 所以站点 i + 1 就是最低点（起点）
			start = i + 1
			minSum = sum
		}
	}
	if sum < 0 {
		// 总油量小于总的消耗，无解
		return -1
	}
	// 环形数组特性
	if start == n {
		return 0
	}
	return start
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=134
