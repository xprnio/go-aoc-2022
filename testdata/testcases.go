package testdata

import "github.com/xprnio/go-aoc-2022/pkg/day01"

type ElvenInput day01.ElvenFoodInventory

type ElvenTestCase struct {
	Input            ElvenInput
	ExpectedCalories []day01.ElvenFoodCalories
}

var BaseElvenTestCase = ElvenTestCase{
	Input: ElvenInput{
		day01.ElvenFoodBasket{1000, 2000, 3000},
		day01.ElvenFoodBasket{4000},
		day01.ElvenFoodBasket{5000, 6000},
		day01.ElvenFoodBasket{7000, 8000, 9000},
		day01.ElvenFoodBasket{10000},
	},
	ExpectedCalories: []day01.ElvenFoodCalories{
		24000,
		11000,
		10000,
		6000,
		4000,
	},
}

var ElvenTestCases = []ElvenTestCase{
	BaseElvenTestCase,
}
