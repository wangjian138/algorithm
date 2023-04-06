package main

func main() {

}

// https://leetcode.cn/problems/trapping-rain-water/
func trap(height []int) int {
	sum := 0
	n := len(height)
	lh := make([]int, n)
	rh := make([]int, n)
	lh[0] = height[0]
	rh[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lh[i] = max(lh[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rh[i] = max(rh[i+1], height[i])
	}
	for i := 1; i < n-1; i++ {
		h := min(rh[i], lh[i]) - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
