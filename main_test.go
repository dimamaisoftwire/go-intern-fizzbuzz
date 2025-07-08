package main

import (
	"fmt"
	"reflect"
	"testing"
)

var ruleTests = []struct {
	number int
	words  []string
}{
	{1, []string{}},
	{6, []string{"Fizz"}},
	{10, []string{"Buzz"}},
	{15, []string{"Fizz", "Buzz"}},
	{7, []string{"Bang"}},
	{21, []string{"Fizz", "Bang"}},
	{35, []string{"Buzz", "Bang"}},
	{11, []string{"Bong"}},
	{33, []string{"Bong"}},
	{13, []string{"Fezz"}},
	{65, []string{"Fezz", "Buzz"}},
	{39, []string{"Fizz", "Fezz"}},
	{255, []string{"Buzz", "Fizz"}},
}

var formatTests = []struct {
	words    []string
	expected string
}{
	{[]string{"Fizz"}, "Fizz"},
	{[]string{"Fizz", "Buzz"}, "FizzBuzz"},
	{[]string{"Fizz", "Buzz", "Bang"}, "FizzBuzzBang"},
	{[]string{}, ""},
}

func TestDetermineOutput(t *testing.T) {
	for _, test := range ruleTests {
		result := determineOutput(test.number)
		if !reflect.DeepEqual(result, test.words) {
			errorMsg := fmt.Sprint("incorrect rule output, for number ", test.number, " got: ", result, " expected: ", test.words)
			t.Error(errorMsg)
		}
	}
}
func TestFormatOutput(t *testing.T) {
	for _, test := range formatTests {
		result := formatOutput(test.words)
		if !(result == test.expected) {
			errorMsg := fmt.Sprint("incorrect formatted output, for words ", test.words, " got: ", result, " expected: ", test.expected)
			t.Error(errorMsg)
		}
	}
}
