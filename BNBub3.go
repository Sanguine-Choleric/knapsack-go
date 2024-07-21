package main

type BNBub3Solver struct {
	best       *KnapsackSolution
	current    *KnapsackSolution
	kp         *KnapsackProblem
	sumTaken   int
	sumUntaken int
}

func (bnb *BNBub3Solver) Solve() {
	bnb.FindSolution(0)
}

func (bnb *BNBub3Solver) FindSolution(itemNum int) {

}

func (bnb *BNBub3Solver) InitialFractionalKnapsack(items []Item) int {
	solution := float32(0)
	for _, e := range items {
		if bnb.kp.capacity <= 0 {
			break
		}

		if e.weight <= bnb.kp.capacity {
			solution += float32(e.value)
			bnb.kp.capacity -= e.weight
		} else {
			// Omitting fractional part for initial solution
			break
		}
	}
	return int(solution + 1)
}