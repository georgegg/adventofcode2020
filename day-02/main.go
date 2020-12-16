package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type policy struct {
	min, max int
	char     string
	pwd      string
}

// Part one:
// Each line gives the password policy and then the password.
// The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid.
// For example, [1-3 a] means that the password must contain a at least 1 time and at most 3 times.
func (p *policy) isValidForMinMax() bool {
	count := strings.Count(p.pwd, p.char)

	return count >= p.min && count <= p.max
}

// Part two:
// Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character,
// and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter.
// Other occurrences of the letter are irrelevant for the purposes of policy enforcement.
func (p *policy) isValidForPositions() bool {
	if p.min-1 < 0 || p.min > len(p.pwd) || p.max-1 < 0 || p.max > len(p.pwd) {
		return false
	}

	c1 := string(p.pwd[p.min-1])
	c2 := string(p.pwd[p.max-1])

	return (c1 == p.char) != (c2 == p.char)
}

func main() {
	input := flag.String("input", "", "--input=/input/file/path")

	flag.Parse()

	if err := run(*input); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(inputFile string) error {
	fmt.Println("--- Day 2: Password Philosophy ---")
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var counterForMinMax, counterForPositions int

	// split by line
	inputs := strings.Split(string(content), "\n")

	for _, s := range inputs {
		policy, err := parceNewPolicy(s)
		if err != nil {
			log.Println(err)
			continue
		}
		if policy != nil && policy.isValidForMinMax() {
			counterForMinMax++
		}
		if policy != nil && policy.isValidForPositions() {
			counterForPositions++
		}
	}
	fmt.Printf("valid passwords for Min/Max policy: %d\n", counterForMinMax)
	fmt.Printf("valid passwords for Positions policy: %d\n", counterForPositions)

	return nil
}

func parceNewPolicy(input string) (*policy, error) {
	if len(input) == 0 {
		return nil, nil
	}
	splitBySpace := strings.Split(input, " ")
	if len(splitBySpace) < 3 {
		return nil, fmt.Errorf("couldn't parce policy, [%s]", input)
	}
	minMax := strings.Split(splitBySpace[0], "-")
	char := strings.Split(splitBySpace[1], ":")[0]
	passwd := splitBySpace[2]

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, fmt.Errorf("couldn't parce policy min, [%s]", input)
	}

	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return nil, fmt.Errorf("couldn't parce policy max, [%s]", input)
	}

	return &policy{
		min:  min,
		max:  max,
		char: char,
		pwd:  passwd,
	}, nil
}
