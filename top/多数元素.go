package main

func main() {

}

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func majorityElement(nums []int) int {
	// 我们想寻找的那个众数
	target := 0
	// 计数器（类比带电粒子例子中的带电性）
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			// 当计数器为 0 时，假设 nums[i] 就是众数
			target = nums[i]
			// 众数出现了一次
			count = 1
		} else if nums[i] == target {
			// 如果遇到的是目标众数，计数器累加
			count++
		} else {
			// 如果遇到的不是目标众数，计数器递减
			count--
		}
	}
	// 回想带电粒子的例子
	// 此时的 count 必然大于 0，此时的 target 必然就是目标众数
	return target
}
