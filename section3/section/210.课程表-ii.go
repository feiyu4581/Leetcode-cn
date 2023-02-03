package section

import "fmt"

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start

type orderPrerequisite struct {
	course int
	next   []*orderPrerequisite
	count  int
	choose bool
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	courses := make(map[int]*orderPrerequisite, numCourses)
	for i := 0; i < numCourses; i++ {
		courses[i] = &orderPrerequisite{course: i}
	}

	for _, prerequisite := range prerequisites {
		courses[prerequisite[0]].count += 1
		courses[prerequisite[1]].next = append(courses[prerequisite[1]].next, courses[prerequisite[0]])
	}

	result := make([]int, 0, numCourses)
	for len(result) < numCourses {
		chose := false
		for i := 0; i < numCourses; i++ {
			if !courses[i].choose && courses[i].count == 0 {
				chose = true
				result = append(result, courses[i].course)
				courses[i].choose = true
				for _, next := range courses[i].next {
					next.count -= 1
				}
			}
		}

		if !chose {
			return []int{}
		}
	}

	return result
}

func Test210() {
	// 0, 1
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	// 0, 2, 1, 3
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	// 0
	fmt.Println(findOrder(1, [][]int{}))
}

// @lc code=end
