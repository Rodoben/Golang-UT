package main

import "testing"

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
		{name: "palindrome3", input: "namann", expected: false, msg: "Not Palindrome!"},
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
