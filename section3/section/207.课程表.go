package section

import "fmt"

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

type coursePrerequisite struct {
	course  int
	next    []*coursePrerequisite
	visited bool
}

func dfsPrerequisite(course *coursePrerequisite, history map[int]struct{}) bool {
	if course.visited {
		return true
	}
	if _, ok := history[course.course]; ok {
		return false
	}

	history[course.course] = struct{}{}
	for _, next := range course.next {
		if !dfsPrerequisite(next, history) {
			return false
		}
	}

	delete(history, course.course)
	course.visited = true
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make(map[int]*coursePrerequisite, numCourses)
	for i := 0; i < numCourses; i++ {
		courses[i] = &coursePrerequisite{course: i}
	}

	for _, prerequisite := range prerequisites {
		courses[prerequisite[0]].next = append(courses[prerequisite[0]].next, courses[prerequisite[1]])
	}

	for turn := 0; turn < numCourses; turn++ {
		if !dfsPrerequisite(courses[turn], make(map[int]struct{}, 0)) {
			return false
		}
	}

	return true
}

func Test207() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}) == false)
	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	fmt.Println(canFinish(4, [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}) == false)
}

// @lc code=end
