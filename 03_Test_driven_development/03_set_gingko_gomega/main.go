package main

import (
	"fmt"
	set "tdd/03_set_gingko_gomega/set"
)

func main() {
	s := set.NewSet()
	fmt.Println(s)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())
	fmt.Println(s.Contains("red"))
	s.Add("red")
	fmt.Println(s.Contains("red"))
	fmt.Println(s.Size())
}
