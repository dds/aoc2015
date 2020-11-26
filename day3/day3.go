package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2015/lib/inputs"
)

var Input = inputs.Day3()

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

type point struct {
	x, y int
}

func (p *point) move(c byte) {
	switch c {
	case '>':
		p.x++
	case '^':
		p.y++
	case '<':
		p.x--
	case 'v':
		p.y--
	}
}

func part1(input string) (rc int) {
	m := map[point]int{}
	p := &point{}

	for _, c := range input {
		m[*p]++
		p.move(byte(c))
	}

	rc = len(m)
	return
}

func part2(input string) (rc int) {
	m := map[point]int{}
	p1, p2 := &point{}, &point{}

	for i := 0; i < len(input); i += 2 {
		m[*p1]++
		m[*p2]++
		p1.move(byte(input[i]))
		p2.move(byte(input[i+1]))
	}

	rc = len(m)
	return
}
