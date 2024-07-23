package main

type BNBub2Solver struct {
	best          *KnapsackSolution
	current       *KnapsackSolution
	kp            *KnapsackProblem
	sumTaken      int
	sumUnTaken    int
	nodesExplored int
}

func (bnb *BNBub2Solver) Solve() {
	bnb.FindSolution(0)
}

func (bnb *BNBub2Solver) FindSolution(itemNum int) {
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

func (bnb *BNBub2Solver) calculateUB(itemNum int) int {
	// Sum taken + sum undecided that fit in remaining capacity
	// Taken: bnb.sumTaken
	// Untaken: bnb.sumUnTaken
	// Undecided: bnb.kp.totalValue - taken - untaken

	// Undecided that fit
	remainingCapacity := bnb.kp.capacity - bnb.current.SumWeights(bnb.kp)
	sumUndecidedFit := 0
	for i := itemNum; i < len(bnb.kp.items); i++ {
		if !bnb.current.takenItems[i] && (bnb.kp.items[i].weight <= remainingCapacity) {
			sumUndecidedFit += bnb.kp.items[i].value
		}
	}

	return sumUndecidedFit + bnb.sumTaken
}

func (bnb *BNBub2Solver) Take(itemNum int) {
	bnb.current.takenItems[itemNum] = true
	bnb.sumTaken += bnb.kp.items[itemNum].value
	bnb.nodesExplored++
}

func (bnb *BNBub2Solver) UndoTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumTaken -= bnb.kp.items[itemNum].value
}

func (bnb *BNBub2Solver) DontTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumUnTaken += bnb.kp.items[itemNum].value
	bnb.nodesExplored++
}

func (bnb *BNBub2Solver) UndoDontTake(itemNum int) {
	bnb.sumUnTaken -= bnb.kp.items[itemNum].value
}
