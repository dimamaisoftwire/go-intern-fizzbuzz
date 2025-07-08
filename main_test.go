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
}

/*var formatTests = []struct {
	input    []string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"", "Ok, no greeting for you"},
}*/

func TestDetermineOutput(t *testing.T) {
	for _, test := range ruleTests {
		result := determineOutput(test.number)
		if !reflect.DeepEqual(result, test.words) {
			errorMsg := fmt.Sprint("incorrect rule output, for number ", test.number, " got: ", result, " expected: ", test.words)
			t.Error(errorMsg)
		}
	}
}
