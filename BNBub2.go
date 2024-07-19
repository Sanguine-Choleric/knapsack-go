package main

type BNBub2Solver struct {
	best       *KnapsackSolution
	current    *KnapsackSolution
	kp         *KnapsackProblem
	sumTaken   int
	sumUnTaken int
}

func (bnb *BNBub2Solver) Solve() {
	bnb.FindSolution(0)
}

func (bnb *BNBub2Solver) FindSolution(itemNum int) {
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
	if (bnb.kp.totalValue() - bnb.current.SumValues(bnb.kp)) <= bnb.best.SumValues(bnb.kp) {
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

func (bnb *BNBub2Solver) calcUB() int {
	return 0
}

func (bnb *BNBub2Solver) Take(itemNum int) {
	bnb.current.takenItems[itemNum] = true
	bnb.sumTaken += bnb.kp.values[itemNum]
}

func (bnb *BNBub2Solver) UndoTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumTaken -= bnb.kp.values[itemNum]
}

func (bnb *BNBub2Solver) DontTake(itemNum int) {
	bnb.current.takenItems[itemNum] = false
	bnb.sumUnTaken += bnb.kp.values[itemNum]
}

func (bnb *BNBub2Solver) UndoDontTake(itemNum int) {
	bnb.sumUnTaken -= bnb.kp.values[itemNum]
}
