package main

func main() {

}

// https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	index := removeElement(nums, 0)
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] == val {
			fast++
			continue
		}
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return slow
}
