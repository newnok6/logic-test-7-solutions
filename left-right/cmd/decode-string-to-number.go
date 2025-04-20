package main

import (
	"fmt"
	"os"
	"strconv"
)

func decodeLetter(s string, value int) (leftValue int, rightValue int) {
	switch s {
	case "L":
		leftValue = value
		rightValue = value - 1
	case "R":
		leftValue = value - 1
		rightValue = value
	case "=":
		leftValue = value
		rightValue = value
	default:
		leftValue = 0
		rightValue = 0
	}
	return
}

func main() {
	// Read pattern from command line
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <pattern> (e.g., LRL=R)")
		return
	}
	s := os.Args[1]

	firstLetter := string(s[0])
	value := [][]int{}
	newString := ""

	switch firstLetter {
	case "L":
		value = append(value, []int{4, 2})
		newString += strconv.Itoa(value[0][0])
	case "R":
		value = append(value, []int{2, 4})
		newString += strconv.Itoa(value[0][1])
	case "=":
		value = append(value, []int{2, 2})
		newString += strconv.Itoa(value[0][0])
	}

	for _, char := range s[1:] {
		lastValue := value[len(value)-1]
		leftValue, rightValue := decodeLetter(string(char), lastValue[1])
		value = append(value, []int{leftValue, rightValue})

		switch string(char) {
		case "L":
			newString += strconv.Itoa(leftValue)
		case "R":
			newString += strconv.Itoa(rightValue)
		case "=":
			newString += strconv.Itoa(leftValue)
		}
	}

	fmt.Println("Decoded value:", value)
	fmt.Println("Decoded string:", newString)
}
