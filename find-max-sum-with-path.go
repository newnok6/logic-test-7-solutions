package main

import (
	"fmt"
)

func maxPathSum(data [][]int) int {
	startRow := len(data) - 2              // Row start on length of the triangle - 1
	for row := startRow; row >= 0; row-- { // calculate from the startRow and traveling to upper node
		startColumn := len(data[row]) // start with index 0 on each row
		for col := 0; col < startColumn; col++ {
			maxValue := 0
			leftValue := data[row+1][col]
			rightValue := data[row+1][col+1]

			// on each row compare between 2 value currIndex and nextIndex
			if leftValue > rightValue {
				maxValue = leftValue
			} else {
				maxValue = rightValue
			}
			data[row][col] += maxValue // set a new value that already sum from lower node
			fmt.Println(data[row][col])
		}
	}
	return data[0][0]
}

func main() {
	data := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	result := maxPathSum(data)
	fmt.Println("Maximum path sum:", result)
}
