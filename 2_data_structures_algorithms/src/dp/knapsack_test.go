package dsa

import (
	"fmt"
	"testing"
)

/*
You are given a set of items, each with a weight and a value, and a knapsack with a maximum capacity.
The goal is to determine the maximum value that can be obtained by selecting a subset of the items to include in the knapsack,
such that the total weight of the selected items does not exceed the knapsack's capacity.
*/
func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weights[i-1]])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func TestKnapsack(t *testing.T) {
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5

	maxVal := knapsack(weights, values, capacity)
	fmt.Println("Maximum value:", maxVal)
}
