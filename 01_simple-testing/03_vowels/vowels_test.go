package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_Vowels(t *testing.T) {
	x := []string{}
	v, err := vowels(x)
	if v != nil {
		t.Errorf("expected nil got %v", v)
	}
	expectedError := fmt.Errorf("%v: Empty Slice", x)
	if errors.Is(err, expectedError) {
		t.Errorf("expected error %v, got %v", expectedError, err)
	}

	y := []string{"a", "a", "a", "e", "1"}
	v, err = vowels(y)
	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}
	expectedMap := map[string]int{"a": 3, "e": 1}
	fmt.Println(v)
	if reflect.DeepEqual(v, expectedMap) {
		fmt.Println("wwwww")
		t.Error("Wrongly passed")
	}
	printMap(v)

}

func TestTable_Vowels(t *testing.T) {
	isVowel := []struct {
		name          string
		input         []string
		expectedMap   map[string]int
		expectedError error
	}{
		{
			name:          "Empty Input",
			input:         []string{},
			expectedMap:   map[string]int{},
			expectedError: errors.New("[]: Empty Slice"),
		},
		{
			name:          "Valid Input",
			input:         []string{"a", "a", "a", "e"},
			expectedMap:   map[string]int{"a": 3, "e": 1, "i": 0, "o": 0, "u": 0},
			expectedError: nil,
		},
	}

	for _, test := range isVowel {

		v, err := vowels(test.input)

		if !reflect.DeepEqual(v, test.expectedMap) {
			t.Errorf("Maps are not equal. Expected: %v, Got: %v", test.expectedMap, v)
		}
		expectedError := fmt.Errorf("%v: Empty Slice", test.input)
		fmt.Println(expectedError)
		if errors.Is(err, expectedError) {
			t.Errorf("expected error %v, got %v", test.expectedError, err)
		}

	}
}
