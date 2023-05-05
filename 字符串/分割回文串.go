package main

import "container/list"

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/*
list包需在solution code中导入：container/list
*/
func partition(s string) [][]string {
	res := make([][]string, 0)
	track := list.New() //linked list

	/*
	   用双指针技巧判断 s[lo..hi] 是否是一个回文串
	*/
	isPalindrome := func(s string, lo int, hi int) bool {
		for lo < hi {
			if s[lo] != s[hi] {
				return false
			}
			lo++
			hi--
		}
		return true
	}

	/*
	   回溯算法框架
	*/
	var backtrack func(s string, start int)
	backtrack = func(s string, start int) {
		if start == len(s) {
			//base case，走到叶子节点
			//即整个s被成功分割为若干个回文子串，记下答案
			tmp := make([]string, 0)
			for val := track.Front(); val != nil; val = val.Next() {
				tmp = append(tmp, val.Value.(string))
			}
			res = append(res, tmp)
		}

		for i := start; i < len(s); i++ {
			if !isPalindrome(s, start, i) {
				//s[start..i]不是回文串，不能分割
				continue
			}

			//s[start..i]是一个回文串，可以进行分割
			//做选择，把 s[start..i] 放入路径列表中
			track.PushBack(s[start : i+1])
			//进入回溯树的下一层，继续切分 s[i+1..]
			backtrack(s, i+1)
			//撤销选择
			track.Remove(track.Back())
		}
	}
	backtrack(s, 0)
	return res
}
