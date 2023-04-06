package main

func main() {

}

// https://leetcode.cn/problems/rotate-image/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先沿对角线反转二维矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 然后反转二维矩阵的每一行
	for _, row := range matrix {
		reverse(row)
	}
}

// 反转一维数组
func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=48
