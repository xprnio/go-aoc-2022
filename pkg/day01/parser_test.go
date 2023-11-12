package day01_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/xprnio/go-aoc-2022/pkg/day01"
)

const (
	TestInput = "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"
)

func assertFoodList(t *testing.T, expected, received day01.ElvenFoodBasket) error {
	if len(expected) != len(received) {
		message := fmt.Sprintf(
			"Expected aoc.ParseInput to return array of %d items, received %d",
			len(expected),
			len(received),
		)
		return errors.New(message)
	}

	return nil
}

func assertElvenFood(t *testing.T, expected, received []day01.ElvenFoodBasket) {
	if len(expected) != len(received) {
		t.Fatalf(
			"Expected aoc.ParseInput to return array of %d items, received %d\n",
			len(expected),
			len(received),
		)
	}

	for index := range received {
		err := assertFoodList(t, expected[index], received[index])
		if err != nil {
			t.Fatalf("Error at index %d:%s\n", index, err)
		}
	}
}

func TestParseInput(t *testing.T) {
	output, err := day01.ParseInput(TestInput)
	expected := []day01.ElvenFoodBasket{
		day01.ElvenFoodBasket{1000, 2000, 3000},
		day01.ElvenFoodBasket{4000},
		day01.ElvenFoodBasket{5000, 6000},
		day01.ElvenFoodBasket{7000, 8000, 9000},
		day01.ElvenFoodBasket{10000},
	}

	if err != nil {
		t.Fatalf("Expected aoc.ParseInput not to return error, received %s\n", err)
	}

	assertElvenFood(t, expected, output)
}
