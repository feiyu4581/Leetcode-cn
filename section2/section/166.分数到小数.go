package section

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 || denominator == 0 {
		return "0"
	}
	nums := make([]uint8, 0)
	cache := make(map[int]int, 0)

	index := 0
	current := 0

	includePoint := -1
	isNegative := 1
	if numerator < 0 {
		isNegative *= -1
		numerator *= -1
	}

	if denominator < 0 {
		isNegative *= -1
		denominator *= -1
	}

	numList := strconv.Itoa(numerator)
	for current != 0 || index < len(numList) {
		current = current * 10
		if index < len(numList) {
			current += int(numList[index] - '0')
		} else if index == len(numList) {
			nums = append(nums, '.')
			includePoint = len(nums) - 1
		}

		if cacheIndex, ok := cache[current]; ok {
			nums = append(nums, ')')
			tempNums := make([]uint8, 0, len(nums)+1)
			tempNums = append(tempNums, nums[0:cacheIndex]...)
			tempNums = append(tempNums, '(')
			tempNums = append(tempNums, nums[cacheIndex:]...)
			nums = tempNums
			break
		}

		if index >= len(numList) {
			cache[current] = len(nums)
		}

		num := 0
		if current >= denominator {
			num = current / denominator
			current = current % denominator
		}

		nums = append(nums, uint8(num)+'0')
		index += 1
	}

	startNum := 0
	if includePoint == -1 {
		includePoint = len(nums) - 1
	}
	for i := 0; i < includePoint-1; i++ {
		if nums[i] == '0' {
			startNum = i + 1
		} else {
			break
		}
	}
	res := strings.Builder{}
	if isNegative == -1 {
		res.WriteByte('-')
	}

	res.Write(nums[startNum:])
	return res.String()
}

func Test166() {
	fmt.Println(fractionToDecimal(1, 2) == "0.5")
	fmt.Println(fractionToDecimal(1, 2) == "0.5")
	fmt.Println(fractionToDecimal(2, 1) == "2")
	fmt.Println(fractionToDecimal(4, 333) == "0.(012)")
	fmt.Println(fractionToDecimal(50, 8) == "6.25")
	fmt.Println(fractionToDecimal(500, 10) == "50")
	fmt.Println(fractionToDecimal(-50, 8) == "-6.25")
	fmt.Println(fractionToDecimal(50, -8) == "-6.25")
	fmt.Println(fractionToDecimal(-50, -8) == "6.25")
}

// @lc code=end
