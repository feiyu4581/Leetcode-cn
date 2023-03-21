package section

import "fmt"

/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	counter := make(map[byte]int, len(secret))
	consistency, similarity := 0, 0
	for i := range secret {
		if i < len(guess) && secret[i] == guess[i] {
			consistency += 1
		} else {
			counter[secret[i]] += 1
		}
	}

	for i := range guess {
		if i < len(secret) && secret[i] == guess[i] {
			continue
		}

		if num, ok := counter[guess[i]]; ok && num > 0 {
			counter[guess[i]] = num - 1
			similarity += 1
		}
	}

	return fmt.Sprintf("%dA%dB", consistency, similarity)
}

// @lc code=end

func Test299() {
	fmt.Println(getHint("1807", "7810") == "1A3B")
	fmt.Println(getHint("1123", "0111") == "1A1B")
}
