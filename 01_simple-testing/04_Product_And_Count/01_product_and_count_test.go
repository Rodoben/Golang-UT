package main

import (
	"fmt"
	"testing"
)

func Test_Product_And_Count(t *testing.T) {
	input := map[string]int{
		"apple":  3,
		"banana": 5,
		"cherry": 2,
		"date":   4,
		"fig":    7,
		"mango":  15,
	}
	p, c, err := calculateProductandCount(input)
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
	if p != 540 {
		t.Errorf("Expected 540, but got %v", p)
	}
	if c != 1 {
		t.Errorf("Expected 1, but got %v", 1)
	}
}

func Test_Table_ProductAndCount(t *testing.T) {
	input1 := map[string]int{
		"apple":  3,
		"banana": 5,
		"cherry": 2,
		"date":   4,
		"fig":    7,
	}
	input2 := map[string]int{
		"apple":  3,
		"banana": 5,
		"cherry": 2,
		"date":   4,
		"fig":    7,
		"mango":  15,
	}

	input3 := map[string]int{
		"apple":  4,
		"banana": 6,
		"cherry": 8,
		"date":   9,
		"fig":    12,
	}
	prodAndcount := []struct {
		name     string
		input    map[string]int
		product  int
		oddcount int
	}{
		{name: "Input 1", input: input1, product: 540, oddcount: 0},
		{name: "Input 2", input: input2, product: 540, oddcount: 1},
		{name: "Input 3", input: input3, product: 1, oddcount: 2},
	}

	for _, test := range prodAndcount {
		p, c, _ := calculateProductandCount(test.input)
		fmt.Println("p:", p, "c:", c)
		if p != test.product || c != test.oddcount {
			t.Errorf("For %s, expected product %v and oddcount %v, but got product %v and oddcount %v",
				test.name, test.product, test.oddcount, p, c)
		}
	}
}
