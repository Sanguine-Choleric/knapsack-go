package main

type KnapsackSolution struct {
	//kp         *KnapsackProblem
	takenItems []bool
}

func (ks *KnapsackSolution) DontTake(itemNum int) {
	ks.takenItems[itemNum] = false
}

func (ks *KnapsackSolution) UndoDontTake(itemNum int) {
	// TODO
}

func (ks *KnapsackSolution) Take(itemNum int) {
	ks.takenItems[itemNum] = true
}

func (ks *KnapsackSolution) UndoTake(itemNum int) {
	ks.takenItems[itemNum] = false
}

func (ks *KnapsackSolution) SumWeights(kp *KnapsackProblem) int {
	sum := 0
	for i, w := range kp.weights {
		if ks.takenItems[i] {
			sum += w
		}
	}
	return sum
}
func (ks *KnapsackSolution) SumValues(kp *KnapsackProblem) int {
	sum := 0
	for i, v := range kp.values {
		if ks.takenItems[i] {
			sum += v
		}
	}
	return sum
}
