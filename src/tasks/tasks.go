package tasks

import (
	"fmt"
)

// PrintTest prints a test string in console
func PrintTest() {
	fmt.Println("Package Tasks loaded!")
}

func SumNumbers(a, b int) int{
	return a + b
}
