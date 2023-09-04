package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_IsPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as a test parameter , got true, but expected false", 0)
	}
	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as a test parameter , got true, but expected false", 7)
	}
	if msg != "7 is a prime number!" {
		t.Error("wrong message returned:", msg)
	}
	result, msg = isPrime(4)
	if result {
		t.Errorf("with %d as a test parameter , got true, but expected false", 4)
	}
	if msg != "4 is not a prime number because it is divisible by 2" {
		t.Error("wrong message returned:", msg)
	}
	result, msg = isPrime(-1)
	if result {
		t.Errorf("with %d as a test parameter , got true, but expected false", -1)
	}
	if msg != "Negative numbers are not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}
}

func Test_IsPrimeTableTest(t *testing.T) {
	isPrimes := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{
			name:     "Prime",
			testNum:  7,
			expected: true,
			msg:      "7 is a prime number!",
		},
		{
			name:     "Not Prime",
			testNum:  4,
			expected: false,
			msg:      "4 is not a prime number because it is divisible by 2",
		},
		{
			name:     "Negative",
			testNum:  -1,
			expected: false,
			msg:      "Negative numbers are not prime, by definition!",
		},
		{
			name:     "Zero",
			testNum:  0,
			expected: false,
			msg:      "0 is not prime, by definition!",
		},
		{
			name:     "One",
			testNum:  1,
			expected: false,
			msg:      "1 is not prime, by definition!",
		},
	}

	for _, e := range isPrimes {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout
	// create a read and write pipe
	r, w, _ := os.Pipe()
	// set os.Stdout to our write pipe
	os.Stdout = w
	prompt()
	// close our writer
	_ = w.Close()
	// reset os.Stdout to what it was before
	os.Stdout = oldOut
	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)
	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
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

	if !strings.Contains(string(out), "Enter a whole number,") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_CheckNumbers1(t *testing.T) {
	input := strings.NewReader("7")
	reader := bufio.NewScanner(input)
	res, _ := checkNumbers(reader)
	if !strings.EqualFold(res, "7 is a prime number!") {
		t.Error("incorrect value returned")
	}
}

func Test_CheckNumbers2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "Zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "One", input: "1", expected: "1 is not prime, by definition!"},
		{name: "Two", input: "2", expected: "2 is a prime number!"},
		{name: "Three", input: "3", expected: "3 is a prime number!"},
		{name: "Negative", input: "-3", expected: "Negative numbers are not prime, by definition!"},
		{name: "Typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.2", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "other", input: "p", expected: "Please enter a whole number!"},
	}
	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected: %s, but got: %s", e.name, e.expected, res)
		}

	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("10\nq\n"))
	go readuserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
