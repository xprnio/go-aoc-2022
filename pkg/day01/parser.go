package day01

import (
	"strconv"
	"strings"
)

func ParseInput(input string) ([]ElvenFoodBasket, error) {
	food := []ElvenFoodBasket{}
	curr := ElvenFoodBasket{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			food = append(food, curr)
			curr = ElvenFoodBasket{}
			continue
		}

		value, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			return nil, err
		}

		calories := ElvenFoodCalories(value)
		curr = append(curr, calories)
	}

	if len(curr) > 0 {
		food = append(food, curr)
	}

	return food, nil
}
