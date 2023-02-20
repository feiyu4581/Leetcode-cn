package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start

func priorityCalculate(s string, index int) (int, int) {
	if index >= len(s) {
		return 0, 0
	}
	numCurrent := make([]byte, 0)

	isMulti := false
	sum := 0
	tryCalculate := func() {
		if len(numCurrent) > 0 {
			num, _ := strconv.Atoi(string(numCurrent))
			numCurrent = make([]byte, 0)
			if isMulti {
				num = -num
				isMulti = false
			}
			sum += num
		}
	}

	for i := index; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numCurrent = append(numCurrent, s[i])
		} else {
			tryCalculate()
			switch s[i] {
			case '-':
				isMulti = true
			case '(':
				nextNum, nextI := priorityCalculate(s, i+1)
				if isMulti {
					nextNum = -nextNum
					isMulti = false
				}
				sum += nextNum
				i = nextI
			case ')':
				tryCalculate()
				return sum, i
			}
		}
	}

	tryCalculate()
	return sum, len(s)
}

func calculate(s string) int {
	// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
	num, _ := priorityCalculate(s, 0)
	return num
}

func Test224() {
	fmt.Println(calculate("1 + 1") == 2)
	fmt.Println(calculate(" 2-1 + 2 ") == 3)
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)") == 23)
	fmt.Println(calculate("1-(     -2)") == 3)
	fmt.Println(calculate("-2+ 1") == -1)
}

// @lc code=end
