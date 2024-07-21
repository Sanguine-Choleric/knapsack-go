package main

type KnapsackProblem struct {
	capacity int
	// weights  []int
	// values   []int
	items    []Item
}

type Item struct {
	value  int
	weight int
}

func (k *KnapsackProblem) totalWeight() int {
	sum := 0
	for i := range k.items {
		sum += k.items[i].weight
	}
	return sum
}

func (k *KnapsackProblem) totalValue() int {
	sum := 0
	for i := range k.items {
		sum += k.items[i].value
	}
	return sum
}
