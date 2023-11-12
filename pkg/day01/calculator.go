package day01

import "sort"

func SumFoodCalories(food ElvenFoodBasket) ElvenFoodCalories {
	sum := ElvenFoodCalories(0)

	for _, cal := range food {
		sum += cal
	}

	return sum
}

func FindTopCalories(input []ElvenFoodBasket) ElvenFoodBasket {
	results := ElvenFoodBasket{}
	for _, food := range input {
		calories := SumFoodCalories(food)
		results = append(results, calories)
	}

	sort.Sort(results)
	results.Reverse()
	return results
}
