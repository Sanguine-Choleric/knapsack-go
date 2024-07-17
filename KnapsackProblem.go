package main

type KnapsackProblem struct {
	capacity int
	weights  []int
	values   []int
}

func (k *KnapsackProblem) totalWeight() {
	sum := 0
	for i := range k.weights {
		sum += k.weights[i]
	}
}

func (k *KnapsackProblem) totalValue() {
	sum := 0
	for i := range k.weights {
		sum += k.values[i]
	}
}
