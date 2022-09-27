package section

import "fmt"

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	l1, l2 := 0, 0
	for l1 < len(version1) || l2 < len(version2) {
		leftVal, rightVal := 0, 0
		if l1 < len(version1) {
			for l1 < len(version1) && version1[l1] != '.' {
				leftVal = leftVal*10 + int(version1[l1]-'0')
				l1 += 1
			}

			if l1 < len(version1) && version1[l1] == '.' {
				l1 += 1
			}
		}

		if l2 < len(version2) {
			for l2 < len(version2) && version2[l2] != '.' {
				rightVal = rightVal*10 + int(version2[l2]-'0')
				l2 += 1
			}

			if l2 < len(version2) && version2[l2] == '.' {
				l2 += 1
			}
		}

		if leftVal > rightVal {
			return 1
		} else if leftVal < rightVal {
			return -1
		}
	}

	return 0
}

func Test165() {
	fmt.Println(compareVersion("1.01", "1.001") == 0)
	fmt.Println(compareVersion("1.0", "1.0.0") == 0)
	fmt.Println(compareVersion("0.1", "1.1") == -1)
}

// @lc code=end
