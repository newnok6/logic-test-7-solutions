package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func maxPathSum(data [][]int) int {
	startRow := len(data) - 2              // Row start on length of the height - 1
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

func readDataFromFile(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var data [][]int
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return data
}

func main() {
	abs, err := filepath.Abs("./find-max-sum-with-path/files/number.json")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
	data := readDataFromFile(abs)

	result := maxPathSum(data)
	fmt.Println("Maximum path sum:", result)
}
