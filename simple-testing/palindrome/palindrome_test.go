package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	result, msg := isPalindrome("naman")
	if !result {
		t.Errorf("with %s as a parameter: expected true got false", "naman")
	}
	if msg != "naman is a Palindrome!" {
		t.Error("Wrong message returned", msg)
	}
	result, msg = isPalindrome("12321")
	if !result {
		t.Errorf("with %s as a parameter: expected true got false", "12321")
	}
	if msg != "12321 is a Palindrome!" {
		t.Error("Wrong message returned", msg)
	}

	result, msg = isPalindrome("namann")
	if result {
		t.Errorf("with %s as a parameter: expected false got false", "namann")
	}
	if msg != "Not Palindrome!" {
		t.Error("Wrong message returned", msg)
	}

	result, msg = isPalindrome("")
	if result {
		t.Errorf("with %s as a parameter: expected false got false", "namann")
	}
	if msg != "Empty string" {
		t.Error("Wrong message returned", msg)
	}
}

func Test_TableTest_isPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		msg      string
	}{
		{name: "palindrome1", input: "malayalam", expected: true, msg: "malayalam is a Palindrome!"},
		{name: "palindrome2", input: "naman", expected: true, msg: "naman is a Palindrome!"},
		{name: "palindrome3", input: "Naman", expected: true, msg: "naman is a Palindrome!"},
		{name: "palindrome4", input: "12321", expected: true, msg: "12321 is a Palindrome!"},
		{name: "palindrome4", input: "123", expected: false, msg: "Not Palindrome!"},
		{name: "palindrome5", input: "", expected: false, msg: "Empty string"},
	}

	for _, test := range tests {
		result, msg := isPalindrome(test.input)

		if test.expected && !result {
			t.Errorf("%s: got false expected true", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s: got true expected false", test.name)
		}

		if test.msg != msg {
			t.Errorf("%s: got: %s but expected: %s", test.name, msg, test.msg)
		}

	}
}

func Test_Prompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if string(out) != "-) " {
		t.Errorf("incorrect prompt: expected -) but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if strings.Contains(string(out), "Enter a Word") {
		t.Errorf("intro text not correct; got %s", string(out))
	}

}

func Test_CheckWord1(t *testing.T) {
	input := strings.NewReader("naman")
	reader := bufio.NewScanner(input)
	res, _ := checkWord(reader)
	if !strings.EqualFold(res, "naman is a Palindrome!") {
		t.Error("incorrect value returned")
	}
}

func Test_CheckWord2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Empty string"},
		{name: "palindrome", input: "naman", expected: "naman is a Palindrome!"},
		{name: "palindrome", input: "12321", expected: "12321 is a Palindrome!"},
		{name: "palindrome", input: "123211", expected: "Not Palindrome!"},
		{name: "quit", input: "q", expected: ""},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)
		res, _ := checkWord(reader)
		if !strings.EqualFold(res, test.expected) {
			t.Errorf("%s: expected: %s, but got: %s", test.name, test.expected, res)
		}
	}
}

func Test_ReadUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("naman\nq\n"))
	go readuserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
