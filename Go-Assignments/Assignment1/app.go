// Assignment 1: create a program to get the prime numbers between 1 to 10 and sum those prime numbers

package main

import (
	"fmt"
)

// Function to check if a number is prime
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0

	fmt.Println("Prime numbers between 1 and 10:")
	for i := 1; i <= 10; i++ {
		if isPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Printf("Sum of prime numbers between 1 and 10: %d\n", sum)
}
