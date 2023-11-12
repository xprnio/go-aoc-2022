package day01_test

import (
	"fmt"
	"testing"

	"github.com/xprnio/go-aoc-2022/pkg/day01"
	"github.com/xprnio/go-aoc-2022/testdata"
)

type CaloriesTestCase struct {
	input  day01.ElvenFoodBasket
	result day01.ElvenFoodCalories
}

var SumFoodCaloriesTestCases = []CaloriesTestCase{
	{
		input:  day01.ElvenFoodBasket{1000, 2000, 3000},
		result: 6000,
	},
	{
		input:  day01.ElvenFoodBasket{4000},
		result: 4000,
	},
	{
		input:  day01.ElvenFoodBasket{5000, 6000},
		result: 11000,
	},
	{
		input:  day01.ElvenFoodBasket{7000, 8000, 9000},
		result: 24000,
	},
	{
		input:  day01.ElvenFoodBasket{10000},
		result: 10000,
	},
}

func TestSumFoodCalories(t *testing.T) {
	for i, test := range SumFoodCaloriesTestCases {
		name := fmt.Sprintf("Test Case %d", i)
		t.Run(name, func(t *testing.T) {
			result := day01.SumFoodCalories(test.input)
			if result != test.result {
				t.Fatalf("Expected aoc.SumFoodCalories to return %d, got %d", test.result, result)
			}
		})
	}
}

func TestFindMostCalories(t *testing.T) {
	for i, test := range testdata.ElvenTestCases {
		name := fmt.Sprintf("Test Case %d", i)
		t.Run(name, func(t *testing.T) {
			input := day01.ElvenFoodInventory(test.Input)
			calories := day01.FindTopCalories(input)
			if len(calories) != len(test.ExpectedCalories) {
				t.Fatalf(
					"Expected aoc.FindMostCalories to return array with %d items, got %d",
					len(test.ExpectedCalories),
					len(calories),
				)
			}

			for index, cal := range calories {
				if cal != test.ExpectedCalories[index] {
					t.Fatalf(
						"Expected calories at index %d to equal %d, got %d",
						index,
						test.ExpectedCalories[index],
						cal,
					)
				}
			}
		})
	}
}
