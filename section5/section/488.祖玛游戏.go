package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=488 lang=golang
 *
 * [488] 祖玛游戏
 */

// @lc code=start

func clearStep(text string) string {
	newText := strings.Builder{}

	start, end := 0, 1
	for end < len(text) {
		if text[end] == text[start] {
			end += 1
		} else if end-start >= 3 {
			start = end
		} else {
			newText.WriteString(text[start:end])
			start = end
		}
	}

	if end-start < 3 {
		newText.WriteString(text[start:end])
	}

	if newText.Len() == 0 {
		return ""
	}

	if ans := newText.String(); len(ans) == len(text) {
		return ans
	} else {
		return clearStep(ans)
	}
}

func findMinStep(board string, hand string) int {
	queues := [][2]string{{board, hand}}
	visits := make(map[[2]string]struct{})

	step := 0
	for len(queues) > 0 {
		newQueues := make([][2]string, 0)
		for _, item := range queues {
			currentBoard, currentHand := item[0], item[1]
			// 广度优先遍历，遍历所有可能的球，以及所有可能的位置
			for i := 0; i < len(currentBoard); i++ {
				for j := 0; j < len(currentHand); j++ {
					// 每个球只会放在连续球的一开始
					if i > 0 && currentBoard[i-1] == currentHand[j] {
						continue
					}

					// 每次拿的球和上次的不一样
					if j > 0 && currentHand[j] == currentHand[j-1] {
						continue
					}

					choose := false
					// 当前球处于连续球的第一个（非第一个的情况已经被第一个条件排除）
					if currentHand[j] == currentBoard[i] {
						choose = true
					}

					// 这是一种特殊的case, [RRWWRRBBRR, WB] ==> [RRWWRRBBRWR, B] ==> ["RRWWRRBBBRWR" ==> "RRWWRRRWR" == "RRWWWR" ==> "RRR" ==> "", ""]
					if i > 0 && currentBoard[i-1] == currentBoard[i] && currentBoard[i] != currentHand[j] {
						choose = true
					}

					if choose {
						newBoard := clearStep(currentBoard[:i] + string(currentHand[j]) + currentBoard[i:])
						if newBoard == "" {
							return step + 1
						}
						newHand := currentHand[:j] + currentHand[j+1:]
						newItem := [2]string{newBoard, newHand}
						// 去除重复情况
						if _, ok := visits[newItem]; !ok {
							newQueues = append(newQueues, newItem)
							visits[newItem] = struct{}{}
						}
					}
				}
			}
		}

		queues = newQueues
		step += 1
	}

	return -1
}

// @lc code=end

func Test488() {
	fmt.Println(findMinStep("RRWWRRBBRR", "WB") == 2)
	fmt.Println(findMinStep("WRRBBW", "RB") == -1)
	fmt.Println(findMinStep("WWRRBBWW", "WRBRW") == 2)
	fmt.Println(findMinStep("G", "GGGGG") == 2)
	fmt.Println(findMinStep("RBYYBBRRB", "YRBGB") == 3)
}
