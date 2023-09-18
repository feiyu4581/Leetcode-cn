package section

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start
func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(nVal)) - 1
	for m := mMax; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)
}

// @lc code=end

func Test483() {
	fmt.Println(smallestGoodBase("13") == "3")
	fmt.Println(smallestGoodBase("4681") == "8")
	fmt.Println(smallestGoodBase("1000000000000000000") == "999999999999999999")
}
