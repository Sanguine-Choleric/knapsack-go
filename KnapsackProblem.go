package main

type KnapsackProblem struct {
	capacity int
	weights  []int
	values   []int
}

func (k *KnapsackProblem) totalWeight() int {
	sum := 0
	for i := range k.weights {
		sum += k.weights[i]
	}
	return sum
}

func (k *KnapsackProblem) totalValue() int {
	sum := 0
	for i := range k.weights {
		sum += k.values[i]
	}
	return sum
}
