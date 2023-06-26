package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	res := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}

	return res
}

// @lc code=end

func Test412() {
	// ["1","2","Fizz"]
	fmt.Println(fizzBuzz(3))
	// ["1","2","Fizz","4","Buzz"]
	fmt.Println(fizzBuzz(5))
	// ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
	fmt.Println(fizzBuzz(15))
}
