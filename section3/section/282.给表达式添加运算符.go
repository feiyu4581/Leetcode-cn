package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=282 lang=golang
 *
 * [282] 给表达式添加运算符
 */

// @lc code=start

type operatorNum struct {
	prefix    []byte
	nums      []int
	operators []byte
}

func copyByte282(source []byte) []byte {
	res := make([]byte, len(source))
	for i := range source {
		res[i] = source[i]
	}
	return res
}

func copyInt282(source []int) []int {
	res := make([]int, len(source))
	for i := range source {
		res[i] = source[i]
	}
	return res
}

func (operator *operatorNum) clone() *operatorNum {
	return &operatorNum{
		prefix:    copyByte282(operator.prefix),
		nums:      copyInt282(operator.nums),
		operators: copyByte282(operator.operators),
	}
}

func (operator *operatorNum) addNum(nums []byte) *operatorNum {
	operator = operator.clone()

	operator.prefix = append(operator.prefix, nums...)
	currentNum, _ := strconv.Atoi(string(nums))
	operator.nums = append(operator.nums, currentNum)
	return operator
}

func (operator *operatorNum) collect() (int, bool) {
	for len(operator.nums) >= 2 && len(operator.operators) >= 1 {
		left, value := operator.nums[len(operator.nums)-2], operator.nums[len(operator.nums)-1]
		operator.nums = operator.nums[:len(operator.nums)-2]

		switch operator.operators[len(operator.operators)-1] {
		case '+':
			operator.nums = append(operator.nums, left+value)
		case '-':
			operator.nums = append(operator.nums, left-value)
		case '*':
			operator.nums = append(operator.nums, left*value)
		}

		operator.operators = operator.operators[:len(operator.operators)-1]
	}

	if len(operator.nums) == 1 && len(operator.operators) == 0 {
		return operator.nums[0], true
	}

	return 0, false
}

func (operator *operatorNum) addOperator(op byte) *operatorNum {
	operator = operator.clone()
	if op != '*' {
		operator.collect()
	}

	operator.operators = append(operator.operators, op)
	operator.prefix = append(operator.prefix, op)
	return operator
}

func addOperators(num string, target int) []string {
	results := make([]string, 0)
	var dfsAddOperators func(index int, operator *operatorNum, needOperator bool)
	dfsAddOperators = func(index int, operator *operatorNum, needOperator bool) {
		if index == len(num) {
			if num, ok := operator.collect(); ok && num == target {
				results = append(results, string(operator.prefix))
			}
			return
		}

		if needOperator {
			dfsAddOperators(index, operator.addOperator('+'), false)
			dfsAddOperators(index, operator.addOperator('-'), false)
			dfsAddOperators(index, operator.addOperator('*'), false)
		} else {
			currentNums := make([]byte, 0)
			for i := index; i < len(num); i++ {
				currentNums = append(currentNums, num[i])
				dfsAddOperators(i+1, operator.addNum(currentNums), true)
				if currentNums[0] == '0' {
					return
				}
			}
		}
	}

	dfsAddOperators(0, &operatorNum{}, false)
	return results
}

// @lc code=end

func Test282() {
	// ["1+2+3", "1*2*3"]
	//fmt.Println(addOperators("123", 6))

	// ["2*3+2", "2+3*2"]
	//fmt.Println(addOperators("232", 8))

	// []
	//fmt.Println(addOperators("3456237490", 9191))

	// ["1*0+5","10-5"]
	fmt.Println(addOperators("105", 5))
}
