package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=335 lang=golang
 *
 * [335] 路径交叉
 */

// @lc code=start

type CrossingMap335 struct {
	columnMap map[int][][2]int
	columns   []int
}

func NewCrossingMap335() *CrossingMap335 {
	return &CrossingMap335{
		columnMap: make(map[int][][2]int, 0),
		columns:   make([]int, 0),
	}
}

func (crossingMap *CrossingMap335) CheckCrossing(x, start, end int, reverse bool) bool {
	i := sort.Search(len(crossingMap.columns), func(i int) bool {
		return crossingMap.columns[i] >= start
	})

	for i < len(crossingMap.columns) {
		num := crossingMap.columns[i]
		if num > end {
			break
		}

		ok := num > start && num <= end
		if reverse {
			ok = num >= start && num < end
		}

		if ok {
			numMaps := crossingMap.columnMap[num]
			j := sort.Search(len(numMaps), func(i int) bool {
				return numMaps[i][1] >= x
			})

			for j < len(numMaps) {
				if numMaps[j][0] > x {
					break
				}

				if numMaps[j][0] <= x && numMaps[j][1] >= x {
					return true
				}

				j += 1
			}
		}

		i += 1
	}

	return false
}

func (crossingMap *CrossingMap335) AddCrossing(x, start, end int) bool {
	if _, ok := crossingMap.columnMap[x]; !ok {
		index := sort.Search(len(crossingMap.columns), func(i int) bool {
			return crossingMap.columns[i] >= x
		})

		crossingMap.columns = append(crossingMap.columns[:index], append([]int{x}, crossingMap.columns[index:]...)...)
		crossingMap.columnMap[x] = [][2]int{{start, end}}
	} else {
		nums := crossingMap.columnMap[x]
		newNums := make([][2]int, 0, len(nums)+1)

		i := sort.Search(len(nums), func(i int) bool {
			return nums[i][1] >= start
		})

		newNums = append(newNums, nums[:i]...)
		for i < len(nums) {
			if nums[i][0] > end {
				newNums = append(newNums, [2]int{start, end})
				newNums = append(newNums, nums[i:]...)
				crossingMap.columnMap[x] = newNums
				return false
			}

			if nums[i][0] <= end && nums[i][1] >= start {
				return true
			} else {
				newNums = append(newNums, nums[i])
			}

			i += 1
		}

		newNums = append(newNums, [2]int{start, end})
		crossingMap.columnMap[x] = newNums
	}
	return false
}

func isSelfCrossing(distance []int) bool {
	if len(distance) >= 100000 {
		return false
	}
	currentX, currentY := 0, 0

	columnMap, rowMap := NewCrossingMap335(), NewCrossingMap335()

	for i, step := range distance {
		switch i % 4 {
		case 0:
			if columnMap.AddCrossing(currentX, currentY, currentY+step) || rowMap.CheckCrossing(currentX, currentY, currentY+step, false) {
				return true
			}
			currentY = currentY + step
		case 1:
			if rowMap.AddCrossing(currentY, currentX-step, currentX) || columnMap.CheckCrossing(currentY, currentX-step, currentX, true) {
				return true
			}
			currentX = currentX - step
		case 2:
			if columnMap.AddCrossing(currentX, currentY-step, currentY) || rowMap.CheckCrossing(currentX, currentY-step, currentY, true) {
				return true
			}
			currentY = currentY - step
		case 3:
			if rowMap.AddCrossing(currentY, currentX, currentX+step) || columnMap.CheckCrossing(currentY, currentX, currentX+step, false) {
				return true
			}
			currentX = currentX + step
		}
	}

	return false
}

// @lc code=end

func Test335() {
	fmt.Println(isSelfCrossing([]int{100, 100, 99, 99, 98, 98, 97, 97, 96, 96, 95, 95, 94, 94, 93, 93, 92, 92, 91, 91, 90, 90, 89, 89, 88, 88, 87, 87, 86, 86, 85, 85, 84, 84, 83, 83, 82, 82, 81, 81, 80, 80, 79, 79, 78, 78, 77, 77, 76, 76, 75, 75, 74, 74, 73, 73, 72, 72, 71, 71, 70, 70, 69, 69, 68, 68, 67, 67, 66, 66, 65, 65, 64, 64, 63, 63, 62, 62, 61, 61, 60, 60, 59, 59, 58, 58, 57, 57, 56, 56, 55, 55, 54, 54, 53, 53, 52, 52, 51, 51, 50, 50, 49, 49, 48, 48, 47, 47, 46, 46, 45, 45, 44, 44, 43, 43, 42, 42, 41, 41, 40, 40, 39, 39, 38, 38, 37, 37, 36, 36, 35, 35, 34, 34, 33, 33, 32, 32, 31, 31, 30, 30, 29, 29, 28, 28, 27, 27, 26, 26, 25, 25, 24, 24, 23, 23, 22, 22, 21, 21, 20, 20, 19, 19, 18, 18, 17, 17, 16, 16, 15, 15, 14, 14, 13, 13, 12, 12, 11, 11, 10, 10, 9, 9, 8, 8, 7, 7, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1, 1, 1, 1, 1}))
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 2, 3, 3, 4, 4, 10, 4, 4, 3, 3, 2, 2, 1, 1}) == false)
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 1, 1}))
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4}) == false)
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 1}))
}
