package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func maxSumPathDFS(arrInput [][]int) int {
	memo := make([][]int, len(arrInput))
	for i := range memo {
		memo[i] = make([]int, len(arrInput[i]))
	}

	return dfs(arrInput, 0, 0, memo)
}

func dfs(arrInput [][]int, row, col int, memo [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}

	if row == len(arrInput)-1 {
		return arrInput[row][col]
	}

	left := dfs(arrInput, row+1, col, memo)
	right := dfs(arrInput, row+1, col+1, memo)

	if left > right {
		memo[row][col] = arrInput[row][col] + left
	} else {
		memo[row][col] = arrInput[row][col] + right
	}

	return memo[row][col]
}

func main() {
	var jsonData = [][]int{}

	data, err := os.ReadFile("hard.json")
	if err != nil {
		log.Panic("Error read file :", err)
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Panic("Error unmarshal :", err)
	}

	if len(jsonData) == 0 {
		fmt.Println(0)
		return
	}

	maxSum := maxSumPathDFS(jsonData)

	fmt.Println("Maximum sum path:", maxSum)
}
