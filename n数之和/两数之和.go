package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int, len(nums))
	for k, v := range nums {
		numsMap[v] = append(numsMap[v], k)
	}
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			if nums[left] == nums[right] {
				return []int{numsMap[nums[left]][0], numsMap[nums[right]][1]}
			}
			return []int{numsMap[nums[left]][0], numsMap[nums[right]][0]}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func twoSum(nums []int, target int) []int {
	// 维护 val -> index 的映射
	valToIndex := make(map[int]int)

	for i, num := range nums {
		// 查表，看看是否有能和 nums[i] 凑出 target 的元素
		need := target - num
		if valToIndex[need] != 0 {
			return []int{valToIndex[need] - 1, i}
		}
		// 存入 val -> index 的映射
		valToIndex[num] = i + 1
	}

	return nil
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=1
