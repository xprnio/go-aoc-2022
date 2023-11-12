package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/xprnio/go-aoc-2022/pkg/day01"
)

func getMaxElements() (uint32, error) {
	if len(os.Args) < 2 {
		return 0, nil
	}

	value, err := strconv.ParseUint(os.Args[1], 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(value), nil
}

func readInput() (string, error) {
	scan := bufio.NewScanner(os.Stdin)
	input := ""

	for scan.Scan() {
		input += scan.Text()
		input += "\n"
	}

	if err := scan.Err(); err != nil {
		return "", nil
	}

	return input, nil
}

func main() {
	max, err := getMaxElements()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid max-n: %s\n", err)
		os.Exit(1)
	}

	input, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %s\n", err)
		os.Exit(1)
	}

	data, err := day01.ParseInput(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %s\n", err)
		os.Exit(1)
	}

	calories := day01.FindTopCalories(data)
	if max > 0 {
		calories = calories[0:max]
	}

	for _, cal := range calories {
		fmt.Fprintf(os.Stdout, "%d\n", cal)
	}
}
