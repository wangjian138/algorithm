package main

func main() {

}

// https://leetcode.cn/problems/course-schedule/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 记录一次 traverse 递归经过的节点
	onPath := make([]bool, numCourses)
	// 记录遍历过的节点，防止走回头路
	visited := make([]bool, numCourses)
	// 记录图中是否有环
	hasCycle := false

	graph := buildGraph(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		// 遍历图中的所有节点
		traverse(graph, i, &hasCycle, visited, onPath)
	}
	// 只要没有循环依赖可以完成所有课程
	return !hasCycle
}

func traverse(graph []LinkedList, s int, hasCycle *bool, visited, onPath []bool) {
	if onPath[s] {
		// 出现环
		*hasCycle = true
	}

	if visited[s] || *hasCycle {
		// 如果已经找到了环，也不用再遍历了
		return
	}
	// 前序遍历代码位置
	visited[s] = true
	onPath[s] = true
	for _, t := range graph[s].list {
		traverse(graph, t, hasCycle, visited, onPath)
	}
	// 后序遍历代码位置
	onPath[s] = false
}

type LinkedList struct {
	list []int
}

func buildGraph(numCourses int, prerequisites [][]int) []LinkedList {
	// 图中共有 numCourses 个节点
	graph := make([]LinkedList, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = LinkedList{list: []int{}}
	}
	for _, edge := range prerequisites {
		from := edge[1]
		to := edge[0]
		// 修完课程 from 才能修课程 to
		// 在图中添加一条从 from 指向 to 的有向边
		graph[from].list = append(graph[from].list, to)
	}
	return graph
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=207
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
