package main

import (
	"fmt"
)

func main() {
	var input1 float32
	var input2 string
	var input3 float32

	fmt.Println("Enter a number:")
	fmt.Scanln(&input1)

	fmt.Println("Enter + or -:")
	fmt.Scanln(&input2)

	fmt.Println("Enter another number:")
	fmt.Scanln(&input3)

	fmt.Println("Result:")
	switch input2 {
	case "+":
		fmt.Println(input1 + input3)
	case "-":
		fmt.Println(input1 - input3)
	default:
		fmt.Println("Unknown operator!")
	}
}
