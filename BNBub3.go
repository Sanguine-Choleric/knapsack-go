package main

import (
	"fmt"
	"sort"
)

type BNBub3Solver struct {
	best    *KnapsackSolution
	current *KnapsackSolution
	kp      *KnapsackProblem
}

func (bnb *BNBub3Solver) Solve() {
	bnb.FindSolution(0)
}

func (bnb *BNBub3Solver) FindSolution(itemNum int) {
	itemCount := len(bnb.kp.weights)

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
	if upperBound <= bnb.best.SumValues(bnb.kp) {
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
	// fractionalValues := make(map[int]float32, len(bnb.kp.values))
	// for i := itemNum; i < len(bnb.kp.values); i++ {
	// 	fractionalValues[i] = float32(bnb.kp.values[i]) / float32(bnb.kp.weights[i])
	// }
	// fmt.Println(fractionalValues)

	type Item struct {
		value  int
		weight int
	}

	organizedValues := make([]Item, len(bnb.kp.values) - itemNum)
	for i := 0; i < len(bnb.kp.values) - itemNum; i++ {
		organizedValues[i].value = bnb.kp.values[i]
		organizedValues[i].weight = bnb.kp.weights[i]
	}

	sort.Slice(organizedValues, func(i, j int) bool { 
		iweightedVal := float32(organizedValues[i].value) / float32(organizedValues[i].weight)
		jweightedVal := float32(organizedValues[j].value) / float32(organizedValues[j].weight)
		return iweightedVal > jweightedVal
	})

	fmt.Println(organizedValues)

	return 100
}

func (bnb *BNBub3Solver) Take(itemNum int) {
	bnb.current.takenItems[itemNum] = true
}

func (bnb *BNBub3Solver) UndoTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
}

func (bnb *BNBub3Solver) DontTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
}

func (bnb *BNBub3Solver) UndoDontTake(itemNum int) {
	// Nothing
}
