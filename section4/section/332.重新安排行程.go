package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start

type node332 struct {
	to      string
	visited bool
}

func findItinerary(tickets [][]string) []string {
	graphs := make(map[string][]*node332, len(tickets))
	for _, ticket := range tickets {
		graphs[ticket[0]] = append(graphs[ticket[0]], &node332{
			to:      ticket[1],
			visited: false,
		})
	}

	for node := range graphs {
		sort.Slice(graphs[node], func(i, j int) bool {
			return graphs[node][i].to <= graphs[node][j].to
		})
	}

	var dfs func(string, int, []string) bool
	dfs = func(node string, index int, paths []string) bool {
		paths[index] = node
		if index == len(paths)-1 {
			return true
		}

		for _, toNode := range graphs[node] {
			if !toNode.visited {
				toNode.visited = true
				if dfs(toNode.to, index+1, paths) {
					return true
				}
				toNode.visited = false
			}
		}

		return false
	}

	paths := make([]string, len(tickets)+1)
	dfs("JFK", 0, paths)
	return paths
}

// @lc code=end

func Test332() {
	// ["JFK","MUC","LHR","SFO","SJC"]
	fmt.Println(findItinerary([][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}))

	// ["JFK","ATL","JFK","SFO","ATL","SFO"]
	fmt.Println(findItinerary([][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}))
}
