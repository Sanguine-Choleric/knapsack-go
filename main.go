package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	return &KnapsackProblem{capacity: sumWeights / 2, weights: weights, values: values}
}

func main() {
	// Problem init
	if len(os.Args) < 1 {
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
	fmt.Println(*k)

	// Solver init
	takenBestItems := make([]bool, itemCount)
	takenCurrItems := make([]bool, itemCount)
	initBest := KnapsackSolution{takenItems: takenBestItems}
	initCurr := KnapsackSolution{takenItems: takenCurrItems}
	//fmt.Println(&initBest)
	//fmt.Println(&initCurr)
	bfSolver := BFSolver{best: &initBest, current: &initCurr, kp: k}
	fmt.Println(bfSolver)
	bfSolver.Solve()
	fmt.Println(bfSolver.best.SumValues(k), ":", bfSolver.best)
}
