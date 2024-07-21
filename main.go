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
	// var weights []int
	// var values []int
	// for i := 0; i < itemCount; i++ {
	// 	weights = append(weights, rand.Intn(100)+1)
	// 	values = append(values, rand.Intn(100)+1)
	// }
	// sumWeights := 0
	// for _, w := range weights {
	// 	sumWeights += w
	// }

	// return &KnapsackProblem{capacity: (sumWeights / 2), weights: weights, values: values}
	items := make([]Item, itemCount)
	sumWeights := 0
	for i := 0; i < itemCount; i++ {
		items[i] = Item{
			value:  rand.Intn(100) + 1,
			weight: rand.Intn(100) + 1,
		}
		sumWeights += items[i].weight
	}
	return &KnapsackProblem{capacity: (sumWeights / 2), items: items}
}

func generateTestKnapsack() *KnapsackProblem {
	// return &KnapsackProblem{capacity: 69, values: []int{47, 61, 95}, weights: []int{57, 37, 49}}
	return &KnapsackProblem{capacity: 41, items: []Item{
		{value: 43, weight: 21},
		{value: 19, weight: 32},
		{value: 60, weight: 29},
	}}
}

func main() {
	// Problem init
	if len(os.Args) < 2 {
		fmt.Println("Need item count as argument")
		return
	}

	// itemCount := os.Args[1]
	itemCount, err := strconv.ParseInt(os.Args[1], 10, 16)
	if err != nil {
		fmt.Println("bad item count")
		return
	}

	k := generateKnapsack(int(itemCount))
	// k := generateTestKnapsack()
	fmt.Println("Cap: ", k.capacity)
	fmt.Println("Items", k.items)

	// Solver init
	takenBestItems := make([]bool, itemCount)
	takenCurrItems := make([]bool, itemCount)
	initBest := KnapsackSolution{takenItems: takenBestItems}
	initCurr := KnapsackSolution{takenItems: takenCurrItems}
	var start time.Time
	//fmt.Println(&initBest)
	//fmt.Println(&initCurr)

	// Brute-Force
	bfSolver := BFSolver{best: &initBest, current: &initCurr, kp: k}
	// fmt.Println(bfSolver)

	start = time.Now()
	bfSolver.Solve()
	bfTime := time.Since(start)
	fmt.Println("BFS\t", bfTime,
		"\t| value =", bfSolver.best.SumValues(k), "| weight =", bfSolver.best.SumWeights(k))
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
	fmt.Println("BT\t", btTime,
		"\t| value =", btSolver.best.SumValues(k), "| weight =", btSolver.best.SumWeights(k))

	// BNB UB1
	takenBestItems = make([]bool, itemCount)
	takenCurrItems = make([]bool, itemCount)
	initBest = KnapsackSolution{takenItems: takenBestItems}
	initCurr = KnapsackSolution{takenItems: takenCurrItems}
	bnb1Solver := BNBub1Solver{best: &initBest, current: &initCurr, kp: k}

	start = time.Now()
	bnb1Solver.Solve()
	bnb1Time := time.Since(start)
	fmt.Println("BNB UB1\t", bnb1Time,
		"\t| value =", bnb1Solver.best.SumValues(k), "| weight =", bnb1Solver.best.SumWeights(k))

	// BNB UB2
	takenBestItems = make([]bool, itemCount)
	takenCurrItems = make([]bool, itemCount)
	initBest = KnapsackSolution{takenItems: takenBestItems}
	initCurr = KnapsackSolution{takenItems: takenCurrItems}
	bnb2Solver := BNBub2Solver{best: &initBest, current: &initCurr, kp: k}

	start = time.Now()
	bnb2Solver.Solve()
	bnb2Time := time.Since(start)
	fmt.Println("BNB UB2\t", bnb2Time,
		"\t| value =", bnb2Solver.best.SumValues(k), "| weight =", bnb2Solver.best.SumWeights(k))

	// BNB UB3 - o(n)
	takenBestItems = make([]bool, itemCount)
	takenCurrItems = make([]bool, itemCount)
	initBest = KnapsackSolution{takenItems: takenBestItems}
	initCurr = KnapsackSolution{takenItems: takenCurrItems}
	bnb3Solver := BNBub3Solver{best: &initBest, current: &initCurr, kp: k}

	start = time.Now()
	bnb3Solver.Solve()
	bnb3Time := time.Since(start)
	fmt.Println("BNB UB3\t", bnb3Time,
		"\t| value =", bnb3Solver.best.SumValues(k), "| weight =", bnb3Solver.best.SumWeights(k))

	// Speedup calcs
	fmt.Println()
	fmt.Println("BF vs BT:\t", (float32(bfTime)-float32(btTime))/float32(btTime)*100, "%")
	fmt.Println("BF vs BNB1:\t", (float32(bfTime)-float32(bnb1Time))/float32(bnb1Time)*100, "%")
	fmt.Println("BF vs BNB2:\t", (float32(bfTime)-float32(bnb2Time))/float32(bnb2Time)*100, "%")
	fmt.Println("BF vs BNB3:\t", (float32(bfTime)-float32(bnb3Time))/float32(bnb3Time)*100, "%")
}
