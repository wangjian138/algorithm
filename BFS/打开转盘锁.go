package main

import "fmt"

func main() {
	//["0201","0101","0102","1212","2002"]
	//	"0202"

	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	findNum := openLock(deadends, "0202")
	fmt.Printf("findNum:%v\n", findNum)
}

// https://leetcode.cn/problems/open-the-lock/
func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool, 0)
	for _, str := range deadends {
		deadendsMap[str] = true
	}

	lockListMap := make(map[string]bool, 0)
	var (
		lockList []string
		lockNum  int
	)
	lockList = append(lockList, "0000")
	lockListMap["0000"] = true

	for len(lockList) > 0 {
		lockListLen := len(lockList)
		for i := 0; i < lockListLen; i++ {
			str := lockList[0]
			lockList = lockList[1:]
			if deadendsMap[str] {
				continue
			}

			if str == target {
				return lockNum
			}

			for j := 0; j < 4; j++ {
				// 新增
				newAddStr := getNewStr(str, j, 1)
				if !lockListMap[newAddStr] {
					lockListMap[newAddStr] = true
					lockList = append(lockList, newAddStr)
				}
				newAddStr = getNewStr(str, j, 2)
				if !lockListMap[newAddStr] {
					lockListMap[newAddStr] = true
					lockList = append(lockList, newAddStr)
				}

			}
		}
		lockNum++
	}

	return -1
}
func getNewStr(str string, num int, opt int) string {
	var newStr []byte
	var numStr byte
	newStr = append(newStr, str[:num]...)

	// 添加
	if opt == 1 {
		numStr = getAddNum(str[num])
	} else {
		numStr = getDesNum(str[num])

	}

	newStr = append(newStr, numStr)
	newStr = append(newStr, str[num+1:]...)

	return string(newStr)
}

func getAddNum(s byte) byte {
	switch s {
	case '9':
		return '0'
	default:
		return s + 1
	}
}

func getDesNum(s byte) byte {
	switch s {
	case '0':
		return '9'
	default:
		return s - 1
	}
}
