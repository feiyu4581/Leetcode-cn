package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start

func removeInvalidParentheses(s string) []string {
	valids := make(map[string]struct{}, 0)
	var dfs func([]byte, int, int)

	clearSlice := func(slice []byte, symbol byte) [][]byte {
		res := make([][]byte, 0)
		for i := range slice {
			if slice[i] == symbol {
				newSlice := make([]byte, len(slice)-1)
				for j := range slice {
					if j < i {
						newSlice[j] = slice[j]
					} else if j > i {
						newSlice[j-1] = slice[j]
					}
				}

				res = append(res, newSlice)
			}
		}

		return res
	}

	var dfsRight func([]byte, []byte, int, int)
	dfsRight = func(expression []byte, prefix []byte, index int, nums int) {
		if index < 0 {
			newPrefix := make([]byte, len(prefix))
			for i := range prefix {
				newPrefix[len(newPrefix)-i-1] = prefix[i]
			}
			valids[string(newPrefix)] = struct{}{}
			return
		}

		prefix = append(prefix, expression[index])
		if expression[index] == ')' {
			dfsRight(expression, prefix, index-1, nums+1)
		} else if expression[index] == '(' {
			if nums == 0 {
				for _, newPrefix := range clearSlice(prefix, '(') {
					dfsRight(expression, newPrefix, index-1, nums)
				}
			} else {
				dfsRight(expression, prefix, index-1, nums-1)
			}
		} else {
			dfsRight(expression, prefix, index-1, nums)
		}
	}

	dfs = func(prefix []byte, index, nums int) {
		if index == len(s) {
			if nums == 0 {
				valids[string(prefix)] = struct{}{}
			} else {
				dfsRight(prefix, make([]byte, 0), len(prefix)-1, 0)
			}

			return
		}

		prefix = append(prefix, s[index])
		if s[index] == '(' {
			dfs(prefix, index+1, nums+1)
		} else if s[index] == ')' {
			if nums == 0 {
				for _, newPrefix := range clearSlice(prefix, ')') {
					dfs(newPrefix, index+1, nums)
				}
			} else {
				dfs(prefix, index+1, nums-1)
			}
		} else {
			dfs(prefix, index+1, nums)
		}
	}

	dfs(make([]byte, 0), 0, 0)

	res := make([]string, 0, len(valids))
	for valid := range valids {
		res = append(res, valid)
	}
	return res
}

// @lc code=end

func Test301() {
	// ["(())()","()()()"]
	fmt.Println(removeInvalidParentheses("()())()"))

	// ["(a())()","(a)()()"]
	fmt.Println(removeInvalidParentheses("(a)())()"))

	// [""]
	fmt.Println(removeInvalidParentheses(")("))
}
