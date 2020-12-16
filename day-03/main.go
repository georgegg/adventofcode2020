package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Slope movement pattern
type Slope struct {
	right, down int
}

// Part2 slopes collection
var Slopes = []Slope{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func main() {
	input := flag.String("input", "", "--input=/input/file/path")
	moves_patterns := flag.Bool("multiple-moves", false, "--multiple-moves to add multiple move patterns")

	flag.Parse()

	if err := run(*input, *moves_patterns); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(inputFile string, moves_patterns bool) error {
	fmt.Println("--- Day 3: Toboggan Trajectory ---")
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// split by line
	lines := strings.Split(string(content), "\n")

	// Slope setup
	slopes := []Slope{
		{3, 1},
	}
	if moves_patterns {
		slopes = Slopes
	}
	counters := make(map[int]int)
	product := 1
	length := len(lines[0])

	for i, slope := range slopes {
		fmt.Printf("Slope movement: right[%d], down[%d]\n", slope.right, slope.down)
		counters[i] = 0
		// Starting point column 0 char 0
		col := 0
		for row := 0; row < len(lines)-1; {
			col += slope.right
			row += slope.down
			if row >= len(lines) || len(string(lines[row])) == 0 {
				break
			}
			if string(lines[row][col%length]) == "#" {
				counters[i]++
			}
		}
		product *= counters[i]
		fmt.Printf("trees encountered: %d\n", counters[i])
	}

	fmt.Printf("trees product of counts per move patterns: %d\n", product)

	return nil
}
