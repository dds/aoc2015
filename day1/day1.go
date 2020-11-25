package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2015/lib/inputs"
)

var Input = inputs.Day1()

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

func part1(input string) int {
	r := 0
	for _, c := range input {
		if c == ')' {
			r--
			continue
		}
		r++
	}
	return r
}

func part2(input string) interface{} {
	r := 0
	for i, c := range input {
		if c == ')' {
			r--
			if r < 0 {
				return i + 1
			}
			continue
		}
		r++
	}
	return -1
}
