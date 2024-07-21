package main

import (
	"fmt"
	"sort"
)

type BNBub3Solver struct {
	best        *KnapsackSolution
	current     *KnapsackSolution
	kp          *KnapsackProblem
	sumTakenV   int
	sumUntakenV int
	sumTakenW   int
}

func (bnb *BNBub3Solver) Solve() {
	bnb.FindInitialBest()
	bnb.FindSolution(0)
}

func (bnb *BNBub3Solver) FindInitialBest() {
	sort.Slice(bnb.kp.items, func(i, j int) bool {
		iFractionalVal := float64(bnb.kp.items[i].value) / float64(bnb.kp.items[i].weight)
		jFractionalVal := float64(bnb.kp.items[j].value) / float64(bnb.kp.items[j].weight)
		return iFractionalVal > jFractionalVal
	})
	fmt.Println(bnb.kp.items)

	initialBest := bnb.InitialFractionalKnapsack()
	fmt.Println(initialBest)

	// Setting best
	sum := 0
	i := 0
	for sum != initialBest {
		bnb.best.takenItems[i] = true
		sum += bnb.kp.items[i].value
		i += 1
	}
	fmt.Println(bnb.best.takenItems)
}

func (bnb *BNBub3Solver) FindSolution(itemNum int) {

}

func (bnb *BNBub3Solver) InitialFractionalKnapsack() int {
	solution := float32(0)
	for _, e := range bnb.kp.items {
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
	return int(solution)
}
