package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2015/lib"
	"github.com/dds/aoc2015/lib/inputs"
)

var Input = lib.InputInts(inputs.Day2(), lib.NumberParser)

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

func paper(l, w, h int) int {
	var e int
	x := 2 * l * w
	y := 2 * w * h
	z := 2 * h * l
	e = x
	if y < e {
		e = y
	}
	if z < e {
		e = z
	}
	return x + y + z + e>>1
}

func part1(input [][]int) int {
	rc := 0
	for _, r := range input {
		// l, w, h
		l, w, h := r[0], r[1], r[2]
		rc += paper(l, w, h)
	}
	return rc
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ribbon(l, w, h int) int {
	r := l * w * h
	m := min(l, min(w, h))
	if l == m {
		return r + 2*l + 2*min(w, h)
	}
	m2 := min(l, max(w, h))
	return r + 2*m + 2*m2
}

func part2(input [][]int) (rc int) {
	for _, r := range input {
		l, w, h := r[0], r[1], r[2]
		rc += ribbon(l, w, h)
	}
	return
}
