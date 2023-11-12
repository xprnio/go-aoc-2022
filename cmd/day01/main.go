package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xprnio/go-aoc-2022/pkg/day01"
)

const inputLineDelimiter = "\n"

func getMaxElements() uint32 {
	max := flag.Uint("max", 0, "Maximum number of elements to return")
	flag.Parse()

	return uint32(*max)
}

func readInput() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", errors.New("No data")
	}

	scan := bufio.NewScanner(os.Stdin)
	input := new(strings.Builder)

	for scan.Scan() {
		input.WriteString(scan.Text())
		input.WriteString(inputLineDelimiter)
	}

	if err := scan.Err(); err != nil {
		return "", nil
	}

	if input.Len() == 0 {
		return "", errors.New("Input is empty")
	}

	return input.String(), nil
}

func main() {
	max := getMaxElements()
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
		fmt.Printf("%d\n", cal)
	}
}
