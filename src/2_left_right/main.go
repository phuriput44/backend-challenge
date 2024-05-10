package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func arrToString(arr []string) string {
	output := ""
	for _, v := range arr {
		output += v
	}
	return output
}

func processInput(input string, ans []string, x int) []string {
	if x > 1 {
		return ans
	}

	left, right := 0, 1

	for _, i := range input {
		if i == 'L' {
			leftVal, _ := strconv.Atoi(ans[left])
			rightVal, _ := strconv.Atoi(ans[right])
			if leftVal <= rightVal {
				rightVal++
				ans[left] = strconv.Itoa(rightVal)
			}
		} else if i == 'R' {
			leftVal, _ := strconv.Atoi(ans[left])
			rightVal, _ := strconv.Atoi(ans[right])
			if rightVal <= leftVal {
				leftVal++
				ans[right] = strconv.Itoa(leftVal)
			}
		} else if i == '=' {
			if ans[left] < ans[right] {
				ans[left] = ans[right]
			} else {
				ans[right] = ans[left]
			}
		}
		left++
		right++
	}

	return processInput(input, ans, x+1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	ans := make([]string, len(input)+1)
	for i := range ans {
		ans[i] = "0"
	}
	x := 0
	finalAns := processInput(input, ans, x)
	fmt.Println(arrToString(finalAns))
}
