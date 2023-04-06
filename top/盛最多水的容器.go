package main

func main() {

}

// https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		// [left, right] 之间的矩形面积
		curArea := min(height[left], height[right]) * (right - left)
		res = max(res, curArea)
		// 双指针技巧，移动较低的一边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// 小函数，求两个数的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 小函数，求两个数的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
