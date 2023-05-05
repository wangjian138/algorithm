package main

func main() {

}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 快慢指针，维护 nums[0..slow] 为结果子数组
	slow, fast := 0, 0
	// 记录一个元素重复的次数
	count := 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2 {
			// 当一个元素重复次数不到 2 次时，也
			slow++
			nums[slow] = nums[fast]
		}
		fast++
		count++
		if fast < len(nums) && nums[fast] != nums[fast-1] {
			// 遇到不同的元素
			count = 0
		}
	}
	// 数组长度为索引 + 1
	return slow + 1
}
