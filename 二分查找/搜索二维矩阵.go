package main

func main() {

}

// https://leetcode.cn/problems/search-a-2d-matrix/?envType=study-plan-v2
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 把二维数组映射到一维
	left, right := 0, m*n-1
	// 前文讲的标准的二分搜索框架
	for left <= right {
		mid := left + (right-left)/2
		if get(matrix, mid) == target {
			return true
		} else if get(matrix, mid) < target {
			left = mid + 1
		} else if get(matrix, mid) > target {
			right = mid - 1
		}
	}
	return false
}

// 通过一维坐标访问二维数组中的元素
func get(matrix [][]int, index int) int {
	n := len(matrix[0])
	// 计算二维中的横纵坐标
	i, j := index/n, index%n
	return matrix[i][j]
}
