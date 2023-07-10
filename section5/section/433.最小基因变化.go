package section

import "fmt"

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	exists := make(map[string]struct{}, len(bank))
	for _, gene := range bank {
		exists[gene] = struct{}{}
	}

	visit := make(map[string]struct{})

	queues := []string{startGene}
	depth := 0

	dirs := []byte{'A', 'C', 'T', 'G'}
	for len(queues) > 0 {
		newQueues := make([]string, 0)
		depth += 1
		for _, gene := range queues {
			for i := range gene {
				for _, dir := range dirs {
					if gene[i] == dir {
						continue
					}

					newGene := gene[:i] + string(dir) + gene[i+1:]
					if _, ok := exists[newGene]; !ok {
						continue
					}

					if newGene == endGene {
						return depth
					}

					if _, ok := visit[newGene]; !ok {
						visit[newGene] = struct{}{}
						newQueues = append(newQueues, newGene)
					}
				}
			}
		}

		queues = newQueues
	}

	return -1
}

// @lc code=end

func Test433() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}) == 1)
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}) == 2)
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}) == 3)
}
