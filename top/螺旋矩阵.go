package main

func main() {

}

// https://leetcode.cn/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	upper_bound, lower_bound := 0, m-1
	left_bound, right_bound := 0, n-1
	res := make([]int, 0, m*n)
	// len(res) == m * n 则遍历完整个数组
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			// 在顶部从左向右遍历
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			// 上边界下移
			upper_bound++
		}

		if left_bound <= right_bound {
			// 在右侧从上向下遍历
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			// 右边界左移
			right_bound--
		}

		if upper_bound <= lower_bound {
			// 在底部从右向左遍历
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			// 下边界上移
			lower_bound--
		}

		if left_bound <= right_bound {
			// 在左侧从下向上遍历
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			// 左边界右移
			left_bound++
		}
	}
	return res
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
