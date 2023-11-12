package day01

type ElvenFoodCalories uint32
type ElvenFoodBasket []ElvenFoodCalories
type ElvenFoodInventory []ElvenFoodBasket

func (list ElvenFoodBasket) Len() int {
	return len(list)
}

func (list ElvenFoodBasket) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list ElvenFoodBasket) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (c ElvenFoodBasket) Reverse() {
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
}
