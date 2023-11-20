package main

import (
	"errors"
	"fmt"
)

func main() {
	inputMap := map[string]int{
		"apple":  4,
		"banana": 6,
		"cherry": 8,
		"date":   9,
		"fig":    12,
	}
	product, count, err := calculateProductandCount(inputMap)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Product of lengths of strings with prime values: %d\n", product)
		fmt.Printf("Count of strings with odd lengths: %d\n", count)
	}
}

func calculateProductandCount(inputmap map[string]int) (int, int, error) {
	product := 1
	oddLengthCount := 0

	for key, value := range inputmap {
		if isPrime(value) {
			product *= len(key)
			fmt.Println(product, "product", len(key), "Key")
		} else {

			if len(key)%2 != 0 {
				oddLengthCount++
			}
		}
	}
	if product == 0 {
		return 0, oddLengthCount, errors.New("product is zero, possibly due to non-prime values in the map")
	}

	if oddLengthCount == 0 {
		return product, 0, errors.New("oddcount is zero")
	}

	return product, oddLengthCount, nil
}

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
