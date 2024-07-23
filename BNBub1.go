package main

type BNBub1Solver struct {
	best          *KnapsackSolution
	current       *KnapsackSolution
	kp            *KnapsackProblem
	sumTaken      int
	sumUntaken    int
	nodesExplored int
}

func (bnb *BNBub1Solver) Solve() {
	bnb.FindSolution(0)
}

func (bnb *BNBub1Solver) FindSolution(itemNum int) {
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
	if (bnb.kp.totalValue() - bnb.sumUntaken) <= bnb.best.SumValues(bnb.kp) {
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

func (bnb *BNBub1Solver) Take(itemNum int) {
	bnb.current.takenItems[itemNum] = true
	bnb.sumTaken += bnb.kp.items[itemNum].value
	bnb.nodesExplored++
}

func (bnb *BNBub1Solver) UndoTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumTaken -= bnb.kp.items[itemNum].value
}

func (bnb *BNBub1Solver) DontTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumUntaken += bnb.kp.items[itemNum].value
	bnb.nodesExplored++
}

func (bnb *BNBub1Solver) UndoDontTake(itemNum int) {
	bnb.sumUntaken -= bnb.kp.items[itemNum].value
}
