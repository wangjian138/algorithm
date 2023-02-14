package main

func main() {

}

// https://leetcode.cn/problems/remove-element/
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
