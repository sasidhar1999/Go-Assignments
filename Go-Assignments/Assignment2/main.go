package main

import (
	"fmt"

	"example.com/calc/calculator"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		result := calculator.Add(num1, num2)
		fmt.Printf("%g + %g = %g\n", num1, num2, result)
	case "-":
		result := calculator.Subtract(num1, num2)
		fmt.Printf("%g - %g = %g\n", num1, num2, result)
	case "*":
		result := calculator.Multiply(num1, num2)
		fmt.Printf("%g * %g = %g\n", num1, num2, result)
	case "/":
		result, err := calculator.Divide(num1, num2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("%g / %g = %g\n", num1, num2, result)
		}
	default:
		fmt.Println("Invalid operator")
	}
}
