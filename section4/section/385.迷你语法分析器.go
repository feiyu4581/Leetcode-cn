package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type IntegerParser struct {
	s       string
	current byte

	index int
}

func (parser *IntegerParser) isEnd() bool {
	return parser.current == 0
}

func (parser *IntegerParser) NextToken() {
	if parser.index >= len(parser.s) {
		parser.current = 0
	} else {
		parser.current = parser.s[parser.index]
		parser.index += 1
	}
}

func (parser *IntegerParser) IsDigit() bool {
	return parser.current == '-' || (parser.current >= '0' && parser.current <= '9')
}

func (parser *IntegerParser) GenerateList() *NestedInteger {
	integerList := &NestedInteger{}

	for parser.current != ']' {
		switch parser.current {
		case '[':
			parser.NextToken()
			integerList.Add(*parser.GenerateList())
		case ']':
			parser.NextToken()
			return integerList
		case ',':
			parser.NextToken()
			continue
		default:
			integerList.Add(*parser.GenerateInteger())
		}
	}

	parser.NextToken()
	return integerList
}

func (parser *IntegerParser) GenerateInteger() *NestedInteger {
	integers := make([]byte, 0)
	flag := 1
	for parser.IsDigit() {
		if parser.current == '-' {
			flag = flag * -1
		} else {
			integers = append(integers, parser.current)
		}

		parser.NextToken()
	}

	integer := &NestedInteger{}
	num, _ := strconv.Atoi(string(integers))
	integer.SetInteger(num * flag)

	return integer
}

func (parser *IntegerParser) Run() *NestedInteger {
	if parser.current == '[' {
		parser.NextToken()
		return parser.GenerateList()
	}

	return parser.GenerateInteger()
}

func deserialize(s string) *NestedInteger {
	parser := IntegerParser{s: s}
	parser.NextToken()

	return parser.Run()
}

// @lc code=end

func Test385() {
	fmt.Println(deserialize("324").String())
	fmt.Println(deserialize("[123,[456,[789]]]").String())
	fmt.Println(deserialize("[123,456,[788,799,833],[[]],10,[]]").String())
}
