package fizzbuzz

import "strconv"

func Run(ints []int) []string {
	str := []string{}
	for _, n := range ints {
		if n == 3 {
			str = append(str, "fizz")
		} else if n == 5 {
			str = append(str, "buzz")

		} else if n == 15 {
			str = append(str, "fizzbuzz")

		} else {
			str = append(str, strconv.Itoa(n))
		}
	}
	return str
}
