package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type KnapsackSolution struct {
	kp         *KnapsackProblem
	takenItems []bool
}

func (ks *KnapsackSolution) sumWeights() int {
	sum := 0
	for i, w := range ks.kp.weights {
		if ks.takenItems[i] {
			sum += w
		}
	}
	return sum
}

func (ks *KnapsackSolution) sumValues() int {
	sum := 0
	for i, v := range ks.kp.values {
		if ks.takenItems[i] {
			sum += v
		}
	}
	return sum
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
}
