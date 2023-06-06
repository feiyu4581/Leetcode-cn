package section

import "fmt"

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start

type Equation399 struct {
	Parents []int
	Weights []float64
}

func NewEquation399(size int) *Equation399 {
	parents := make([]int, size)
	weights := make([]float64, size)

	for i := 0; i < size; i++ {
		parents[i] = i
		weights[i] = 1
	}

	return &Equation399{
		Parents: parents,
		Weights: weights,
	}
}

func (e *Equation399) Find(x int) int {
	if x != e.Parents[x] {
		parent := e.Parents[x]
		e.Parents[x] = e.Find(parent)
		e.Weights[x] = e.Weights[parent] * e.Weights[x]
	}

	return e.Parents[x]
}

func (e *Equation399) Union(x, y int, weight float64) {
	parentX := e.Find(x)
	parentY := e.Find(y)

	if parentX == parentY {
		return
	}

	e.Parents[parentX] = parentY
	e.Weights[parentX] = e.Weights[y] * weight / e.Weights[x]
}

func (e *Equation399) IsConnected(x, y int) float64 {
	parentX := e.Find(x)
	parentY := e.Find(y)

	if parentX == parentY {
		return e.Weights[x] / e.Weights[y]
	}

	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationMaps := make(map[string]int, len(equations)*2)

	e := NewEquation399(len(equations) * 2)
	index := 0
	for i, equation := range equations {
		if _, ok := equationMaps[equation[0]]; !ok {
			equationMaps[equation[0]] = index
			index += 1
		}

		if _, ok := equationMaps[equation[1]]; !ok {
			equationMaps[equation[1]] = index
			index += 1
		}

		e.Union(equationMaps[equation[0]], equationMaps[equation[1]], values[i])
	}

	res := make([]float64, 0, len(queries))
	for _, query := range queries {
		if _, ok := equationMaps[query[0]]; !ok {
			res = append(res, -1)
			continue
		}

		if _, ok := equationMaps[query[1]]; !ok {
			res = append(res, -1)
			continue
		}

		res = append(res, e.IsConnected(equationMaps[query[0]], equationMaps[query[1]]))
	}

	return res
}

// @lc code=end

func Test399() {
	// [6.00000,0.50000,-1.00000,1.00000,-1.00000]
	fmt.Println(calcEquation([][]string{
		{"a", "b"},
		{"b", "c"},
	}, []float64{2.0, 3.0}, [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}))

	// [3.75000,0.40000,5.00000,0.20000]
	fmt.Println(calcEquation([][]string{
		{"a", "b"},
		{"b", "c"},
		{"bc", "cd"},
	}, []float64{1.5, 2.5, 5.0}, [][]string{
		{"a", "c"},
		{"c", "b"},
		{"bc", "cd"},
		{"cd", "bc"},
	}))

	// [0.50000,2.00000,-1.00000,-1.00000]
	fmt.Println(calcEquation([][]string{
		{"a", "b"},
	}, []float64{0.5}, [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "c"},
		{"x", "y"},
	}))
}
