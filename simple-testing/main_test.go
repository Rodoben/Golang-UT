package main

import "testing"

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
