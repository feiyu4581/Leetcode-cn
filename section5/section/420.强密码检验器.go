package section

import "fmt"

/*
 * @lc app=leetcode.cn id=420 lang=golang
 *
 * [420] 强密码检验器
 */

// @lc code=start

func max420(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func strongPasswordChecker(password string) int {
	hasDigit, hasUpper, hasLowwer := 1, 1, 1
	repeatedCount := make([]int, 0)
	count := 0
	for i := 0; i < len(password); i++ {
		if password[i] >= '0' && password[i] <= '9' {
			hasDigit = 0
		} else if password[i] >= 'a' && password[i] <= 'z' {
			hasLowwer = 0
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			hasUpper = 0
		}

		if i == 0 || password[i] == password[i-1] {
			count += 1
		} else {
			if count >= 3 {
				repeatedCount = append(repeatedCount, count)
			}
			count = 1
		}
	}

	if count >= 3 {
		repeatedCount = append(repeatedCount, count)
	}

	errorCount := hasDigit + hasUpper + hasLowwer
	if len(password) < 6 {
		res := 6 - len(password)
		errorCount = max420(0, errorCount-res)
		return res + errorCount
	} else if len(password) <= 20 {
		changeCount := 0
		for _, count := range repeatedCount {
			changeCount += count / 3
		}

		errorCount = max420(0, errorCount-changeCount)
		return changeCount + errorCount
	} else {
		res := len(password) - 20
		for res > 0 {
			changed := false
			for step := 1; step <= 3; step++ {
				for index, count := range repeatedCount {
					if reduceCount := count%3 + 1; count >= 3 && reduceCount <= res && reduceCount == step {
						res -= reduceCount
						repeatedCount[index] -= reduceCount
						changed = true
					}
				}
			}

			if !changed {
				break
			}
		}

		changeCount := 0
		for _, count := range repeatedCount {
			if count >= 3 {
				changeCount += count / 3
			}
		}

		errorCount = max420(0, errorCount-changeCount)
		return errorCount + changeCount + len(password) - 20
	}
}

// @lc code=end

func Test420() {
	fmt.Println(strongPasswordChecker("QQQQQ") == 2)
	fmt.Println(strongPasswordChecker("aaaaAAAAAA000000123456") == 5)
	fmt.Println(strongPasswordChecker("bbaaaaaaaaaaaaaaacccccc") == 8)
	fmt.Println(strongPasswordChecker("a") == 5)
	fmt.Println(strongPasswordChecker("aA1") == 3)
	fmt.Println(strongPasswordChecker("1337C0d3") == 0)
	fmt.Println(strongPasswordChecker("aaa111") == 2)
	fmt.Println(strongPasswordChecker("aaaB1") == 1)
}
