package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type KnapsackSolver interface {
	Solve()
}

func generateKnapsack(itemCount int) *KnapsackProblem {
	var weights []int
	var values []int
	for i := 0; i < itemCount; i++ {
		weights = append(weights, rand.Intn(100)+1)
		values = append(values, rand.Intn(100)+1)
	}
	sumWeights := 0
	for _, w := range weights {
		sumWeights += w
	}

	return &KnapsackProblem{capacity: (sumWeights / 2), weights: weights, values: values}
}

func main() {
	// Problem init
	if len(os.Args) < 2 {
		fmt.Println("Need item count as argument")
		return
	}

	// itemCount := os.Args[1]
	itemCount, err := strconv.ParseInt(os.Args[1], 10, 8)
	if err != nil {
		fmt.Println("bad item count")
		return
	}

	k := generateKnapsack(int(itemCount))
	fmt.Println("Cap: ", k.capacity)
	fmt.Println("Values", k.values)
	fmt.Println("Weight", k.weights)

	// Solver init
	takenBestItems := make([]bool, itemCount)
	takenCurrItems := make([]bool, itemCount)
	initBest := KnapsackSolution{takenItems: takenBestItems}
	initCurr := KnapsackSolution{takenItems: takenCurrItems}
	//fmt.Println(&initBest)
	//fmt.Println(&initCurr)

	// Brute-Force
	bfSolver := BFSolver{best: &initBest, current: &initCurr, kp: k}
	// fmt.Println(bfSolver)

	start := time.Now()
	bfSolver.Solve()
	bfTime := time.Since(start)
	fmt.Println("BFS took", bfTime,
		"| optimal solution is value", bfSolver.best.SumValues(k), "at weight", bfSolver.best.SumWeights(k))
	// fmt.Println(bfSolver.best.SumWeights(k), ":", bfSolver.best.SumValues(k), ":", bfSolver.best)


	// Backtracking
	takenBestItems = make([]bool, itemCount)
	takenCurrItems = make([]bool, itemCount)
	initBest = KnapsackSolution{takenItems: takenBestItems}
	initCurr = KnapsackSolution{takenItems: takenCurrItems}
	btSolver := BTSolver{best: &initBest, current: &initCurr, kp: k}

	start = time.Now()
	btSolver.Solve()
	btTime := time.Since(start)
	fmt.Println("BT took", btTime,
		"| optimal solution is value", btSolver.best.SumValues(k), "at weight", btSolver.best.SumWeights(k))

	// Speedup calcs
	fmt.Println("BF vs BT:", (float32(bfTime) - float32(btTime)) / float32(btTime) * 100, "%")
}
