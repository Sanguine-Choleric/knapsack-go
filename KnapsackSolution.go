package main

type KnapsackSolution struct {
	takenItems []bool
}

func (ks *KnapsackSolution) SumWeights(kp *KnapsackProblem) int {
	sum := 0
	for i, w := range kp.items {
		if ks.takenItems[i] {
			sum += w.weight
		}
	}
	return sum
}
func (ks *KnapsackSolution) SumValues(kp *KnapsackProblem) int {
	sum := 0
	for i, v := range kp.items {
		if ks.takenItems[i] {
			sum += v.value
		}
	}
	return sum
}
