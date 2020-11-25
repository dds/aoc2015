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

func part1(input [][]int) interface{} {
	return input
}

func part2(input [][]int) interface{} {
	return ""
}
