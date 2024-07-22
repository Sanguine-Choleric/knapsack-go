package main

type BTSolver struct {
	best    *KnapsackSolution
	current *KnapsackSolution
	kp      *KnapsackProblem
	takenW  int
}

func (bt *BTSolver) Solve() {
	bt.FindSolution(0)
}

func (bt *BTSolver) FindSolution(itemNum int) {
	itemCount := len(bt.kp.items)

	// currentWeight := bt.current.SumWeights(bt.kp)
	if bt.takenW > bt.kp.capacity {
		// bt.current.DontTake(itemNum - 1)
		return
	}

	// Base Case
	if itemNum == itemCount {
		if bt.takenW <= bt.kp.capacity {
			curr := bt.current.SumValues(bt.kp)
			best := bt.best.SumValues(bt.kp)
			if curr > best {
				//bf.best := bf.current
				copy(bt.best.takenItems, bt.current.takenItems)
			}
		}
		return
	}

	// Take item
	//fmt.Println("Taking", itemNum)
	bt.Take(itemNum)
	bt.FindSolution(itemNum + 1)
	bt.UndoTake(itemNum)
	// No take item
	//fmt.Println("Not Taking", itemNum)
	bt.DontTake(itemNum)
	bt.FindSolution(itemNum + 1)
	bt.UndoDontTake(itemNum)
}

func (bt *BTSolver) Take(itemNum int) {
	bt.current.takenItems[itemNum] = true
	bt.takenW += bt.kp.items[itemNum].weight
}

func (bt *BTSolver) UndoTake(itemNum int) {
	bt.current.takenItems[itemNum] = false
	bt.takenW -= bt.kp.items[itemNum].weight
}

func (bt *BTSolver) DontTake(itemNum int) {
	bt.current.takenItems[itemNum] = false
}

func (bt *BTSolver) UndoDontTake(itemNum int) {
	// Nothing...
}
