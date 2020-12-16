package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := flag.String("input", "", "--input=/input/file/path")

	flag.Parse()

	if err := run(*input); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
func run(inputFile string) error {
	fmt.Println("--- Day 1: Report Repair ---")
	numbers, err := readNums(inputFile)
	if err != nil {
		return err
	}
	// Part one:
	// Find the two entries that sum to 2020; what do you get if you multiply them together?
	result, err := findSumOfTwo(numbers, 2020)
	if err != nil {
		return err
	}

	fmt.Println("found product of pair with sum 2020 =", result)

	// Part two:
	// what is the product of the three entries that sum to 2020?
	result2, err := findSumOfThree(numbers, 2020)
	if err != nil {
		return err
	}
	fmt.Println("found product of triple with sum 2020 =", result2)

	return nil
}

func findSumOfThree(nums []int, sum int) (int, error) {
	for i, n1 := range nums {
		if n1 > 2020 {
			continue
		}
		for j, n2 := range nums[i+1:] {
			if n1+n2 > 2020 {
				continue
			}
			for _, n3 := range nums[j+1:] {
				if n1+n2+n3 == 2020 {
					fmt.Printf("found sum [%d] = [%d] + [%d] + [%d]\n", sum, n1, n2, n3)

					return n1 * n2 * n3, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("Couldn't find triple of nums to sum: [%d]", sum)
}

func findSumOfTwo(nums []int, sum int) (int, error) {
	for i, n1 := range nums {
		for _, n2 := range nums[i+1:] {
			if n1+n2 == sum {
				fmt.Printf("found sum [%d] = [%d] + [%d]\n", sum, n1, n2)

				return n1 * n2, nil
			}
		}
	}

	return 0, fmt.Errorf("Couldn't find pair of nums to sum: [%d]", sum)
}

func readNums(inputFile string) ([]int, error) {
	numbers := make([]int, 0)

	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s, %s", inputFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
