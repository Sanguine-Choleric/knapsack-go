package main

type BFSolver struct {
	best          *KnapsackSolution
	current       *KnapsackSolution
	kp            *KnapsackProblem
	nodesExplored int
}

func (bf *BFSolver) Solve() {
	bf.FindSolution(0)
}

func (bf *BFSolver) FindSolution(itemNum int) {
	itemCount := len(bf.kp.items)

	// Base Case
	if itemNum == itemCount {
		if bf.current.SumWeights(bf.kp) <= bf.kp.capacity {
			curr := bf.current.SumValues(bf.kp)
			best := bf.best.SumValues(bf.kp)
			if curr > best {
				//bf.best := bf.current
				copy(bf.best.takenItems, bf.current.takenItems)
			}
		}
		return
	}
	// Take item
	//fmt.Println("Taking", itemNum)
	bf.Take(itemNum)
	bf.FindSolution(itemNum + 1)
	// No take item
	//fmt.Println("Not Taking", itemNum)
	bf.DontTake(itemNum)
	bf.FindSolution(itemNum + 1)
}

func (bf *BFSolver) Take(itemNum int) {
	bf.current.takenItems[itemNum] = true
	bf.nodesExplored++
}

func (bf *BFSolver) DontTake(itemNum int) {
	bf.current.takenItems[itemNum] = false
	bf.nodesExplored++
}
