package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	intro()
	doneChan := make(chan bool)
	go readuserInput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)
	fmt.Println("Good Bye!")
}

func readuserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkWord(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkWord(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	_, msg := isPalindrome(scanner.Text())
	return msg, false

}

func intro() {
	fmt.Println("Is it Palindrome?")
	fmt.Println("------------")
	fmt.Println("Enter a word, and we'll tell you if it is a palindrome or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPalindrome(s string) (bool, string) {
	if strings.TrimSpace(s) == "" {
		return false, "Empty string"
	}
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false, "Not Palindrome!"
		}
	}
	return true, fmt.Sprintf("%s is a Palindrome!", s)
}
