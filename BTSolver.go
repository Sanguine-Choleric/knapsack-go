package main

type BTSolver struct {
	best *KnapsackSolution
	current *KnapsackSolution
	kp *KnapsackProblem
}

func (bt *BTSolver) Solve() {
	bt.FindSolution(0)
}

func (bt *BTSolver) FindSolution(itemNum int) {
	itemCount := len(bt.kp.items)

	currentWeight := bt.current.SumWeights(bt.kp)
	if currentWeight > bt.kp.capacity {
		// bt.current.DontTake(itemNum - 1)
		return
	}

	// Base Case
	if itemNum == itemCount {
		if currentWeight <= bt.kp.capacity {
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
	bt.current.Take(itemNum)
	bt.FindSolution(itemNum + 1)
	bt.current.UndoTake(itemNum)
	// No take item
	//fmt.Println("Not Taking", itemNum)
	bt.current.DontTake(itemNum)
	bt.FindSolution(itemNum + 1)
	bt.current.UndoDontTake(itemNum)
}