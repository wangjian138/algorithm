package main

func main() {

}

// https://leetcode.cn/problems/contains-duplicate-ii/?envType=study-plan-v2&id=top-interview-150
func containsNearbyDuplicate(nums []int, k int) bool {
	sMap := make(map[int]int, 0)

	for index, n := range nums {
		v, res := sMap[n]
		if !res {
			sMap[n] = index
			continue
		}
		if index-v <= k {
			return true
		}
		sMap[n] = index
	}

	return false
}
