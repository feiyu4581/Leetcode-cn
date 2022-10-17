package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start

var dnaMaps = map[uint8]int{
	'A': 1,
	'C': 2,
	'G': 3,
	'T': 4,
}

var remainderNumber = int(math.Pow(10, 9))

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}

	current := 0
	for i := 0; i < 10; i++ {
		current = current*10 + dnaMaps[s[i]]
	}

	dnaIntMaps := map[int]int{current: 0}
	res := make(map[int]int, 0)
	for i := 1; i+10 <= len(s); i++ {
		current = (current%remainderNumber)*10 + dnaMaps[s[i+9]]
		if index, ok := dnaIntMaps[current]; ok {
			res[current] = index
		} else {
			dnaIntMaps[current] = i
		}
	}

	sequences := make([]string, 0, len(res))
	for _, index := range res {
		sequence := make([]uint8, 0, 10)
		for i := 0; i < 10; i++ {
			sequence = append(sequence, s[index+i])
		}

		sequences = append(sequences, string(sequence))
	}

	return sequences
}

func Test187() {
	// ["AAAAACCCCC","CCCCCAAAAA"]
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	// ["AAAAAAAAAA"]
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}

// @lc code=end
