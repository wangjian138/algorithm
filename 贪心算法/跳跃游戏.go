package main

func main() {

}

// https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	cover := 0
	n := len(nums) - 1
	for i := 0; i <= cover; i++ { // 每次与覆盖值比较
		cover = max(i+nums[i], cover) //每走一步都将 cover 更新为最大值
		if cover >= n {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=55
func canJump1(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		// 不断计算能跳到的最远距离
		farthest = max(farthest, i+nums[i])
		// 可能碰到了 0，卡住跳不动了
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}
