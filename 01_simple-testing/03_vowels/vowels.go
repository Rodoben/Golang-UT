package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	v := []string{"a", "z", "a", "e", "f", "i", "A", "1"}
	//v := []string{}
	m, err := vowels(v)
	if err != nil {
		log.Println(err)
	}
	printMap(m)

}

func vowels(x []string) (map[string]int, error) {

	if len(x) == 0 {
		return map[string]int{}, fmt.Errorf("%v: Empty Slice", x)
	}
	m := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}
	for _, v := range x {
		v = strings.ToLower(v)
		if _, ok := m[v]; ok {
			m[v]++
		}
	}
	return m, nil
}

func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}
