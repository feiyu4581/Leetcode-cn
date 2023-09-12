package section

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start

func tryIPv4Address(queryIp string) bool {
	blocks := strings.Split(queryIp, ".")
	if len(blocks) != 4 {
		return false
	}

	for _, block := range blocks {
		if len(block) == 0 || (len(block) > 1 && block[0] == '0') {
			return false
		}

		num, err := strconv.Atoi(block)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func tryIPv6Address(queryIp string) bool {
	blocks := strings.Split(queryIp, ":")
	if len(blocks) != 8 {
		return false
	}

	for _, block := range blocks {
		if len(block) == 0 || len(block) > 4 {
			return false
		}

		for i := 0; i < len(block); i++ {
			if block[i] >= '0' && block[i] <= '9' {
				continue
			}

			if block[i] >= 'a' && block[i] <= 'f' {
				continue
			}

			if block[i] >= 'A' && block[i] <= 'F' {
				continue
			}

			return false
		}
	}

	return true
}

func validIPAddress(queryIP string) string {
	if tryIPv4Address(queryIP) {
		return "IPv4"
	}

	if tryIPv6Address(queryIP) {
		return "IPv6"
	}

	return "Neither"
}

// @lc code=end

func Test468() {
	fmt.Println(validIPAddress("172.16.254.1") == "IPv4")
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334") == "IPv6")
	fmt.Println(validIPAddress("256.256.256.256") == "Neither")
}
