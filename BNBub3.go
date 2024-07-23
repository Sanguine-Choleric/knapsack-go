package main

import (
	"fmt"
	"sort"
)

type BNBub3Solver struct {
	best          *KnapsackSolution
	current       *KnapsackSolution
	kp            *KnapsackProblem
	sumTakenV     int
	sumUntakenV   int
	sumTakenW     int
	nodesExplored int
}

func (bnb *BNBub3Solver) Solve() {
	sort.Slice(bnb.kp.items, func(i, j int) bool {
		iFractionalVal := float64(bnb.kp.items[i].value) / float64(bnb.kp.items[i].weight)
		jFractionalVal := float64(bnb.kp.items[j].value) / float64(bnb.kp.items[j].weight)
		return iFractionalVal > jFractionalVal
	})
	// bnb.FindInitialBest()
	bnb.FindSolution(0)
}

func (bnb *BNBub3Solver) FindInitialBest() {
	// sort.Slice(bnb.kp.items, func(i, j int) bool {
	// 	iFractionalVal := float64(bnb.kp.items[i].value) / float64(bnb.kp.items[i].weight)
	// 	jFractionalVal := float64(bnb.kp.items[j].value) / float64(bnb.kp.items[j].weight)
	// 	return iFractionalVal > jFractionalVal
	// })
	fmt.Println(bnb.kp.items)

	initialBest := bnb.InitialFractionalKnapsack()
	fmt.Println("Initial:", initialBest)

	// Setting best
	sum := 0
	i := 0
	for sum != initialBest {
		bnb.best.takenItems[i] = true
		sum += bnb.kp.items[i].value
		i += 1
	}
	fmt.Println(bnb.best.takenItems, bnb.best.SumValues(bnb.kp))
}

func (bnb *BNBub3Solver) FindSolution(itemNum int) {
	itemCount := len(bnb.kp.items)

	currentWeight := bnb.current.SumWeights(bnb.kp)
	if currentWeight > bnb.kp.capacity {
		// bnb.current.DontTake(itemNum - 1)
		return
	}

	// Base Case
	if itemNum == itemCount {
		if currentWeight <= bnb.kp.capacity {
			curr := bnb.current.SumValues(bnb.kp)
			best := bnb.best.SumValues(bnb.kp)
			if curr > best {
				//bf.best := bf.current
				copy(bnb.best.takenItems, bnb.current.takenItems)
			}
		}
		return
	}

	// BNB case
	upperBound := bnb.calculateUB(itemNum)
	if upperBound+bnb.sumTakenV <= bnb.best.SumValues(bnb.kp) {
		return
	}

	// Take item
	//fmt.Println("Taking", itemNum)
	bnb.Take(itemNum)
	bnb.FindSolution(itemNum + 1)
	bnb.UndoTake(itemNum)
	// No take item
	//fmt.Println("Not Taking", itemNum)
	bnb.DontTake(itemNum)
	bnb.FindSolution(itemNum + 1)
	bnb.UndoDontTake(itemNum)
}

func (bnb *BNBub3Solver) calculateUB(itemNum int) int {
	// Fractional Knapsack solution

	return bnb.FractionalKnapsack(itemNum)
}

// TODO: Actual implementation
func (bnb *BNBub3Solver) FractionalKnapsack(itemNum int) int {
	remainingCapacity := bnb.kp.capacity - bnb.sumTakenW
	solution := 0.0
	// for _, e := range bnb.kp.items {
	// 	if remainingCapacity <= 0 {
	// 		break
	// 	}

	// 	if e.weight <= remainingCapacity {
	// 		solution += float64(e.weight)
	// 		remainingCapacity -= e.weight
	// 	} else {
	// 		solution += (float64(e.value) * float64(remainingCapacity)) / float64(e.weight)
	// 		break
	// 	}
	// }
	for i := itemNum; i < len(bnb.kp.items); i++ {
		if remainingCapacity <= 0 {
			break
		}

		if bnb.kp.items[i].weight <= remainingCapacity {
			solution += float64(bnb.kp.items[i].value)
			remainingCapacity -= bnb.kp.items[i].weight
		} else {
			solution += float64(bnb.kp.items[i].value) / float64(bnb.kp.items[i].weight) * float64(remainingCapacity)
			break
		}
	}
	return int(solution + 1)
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

func (bnb *BNBub3Solver) Take(itemNum int) {
	bnb.current.takenItems[itemNum] = true
	bnb.sumTakenV += bnb.kp.items[itemNum].value
	bnb.sumTakenW += bnb.kp.items[itemNum].weight
	bnb.nodesExplored++
}

func (bnb *BNBub3Solver) UndoTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumTakenV -= bnb.kp.items[itemNum].value
	bnb.sumTakenW -= bnb.kp.items[itemNum].weight
}

func (bnb *BNBub3Solver) DontTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumUntakenV += bnb.kp.items[itemNum].value
	bnb.nodesExplored++
}

func (bnb *BNBub3Solver) UndoDontTake(itemNum int) {
	bnb.sumUntakenV -= bnb.kp.items[itemNum].value
}
