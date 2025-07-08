package main

import (
	"fmt"
	"strings"
)

var setOfRules map[int]func([]string) []string
var orderOfRules []int

func determineOutput(number int) []string {

	setOfRules = make(map[int]func([]string) []string)
	orderOfRules = make([]int, 0)
	output := make([]string, 0)
	// Go maps do not preserve order of insertion so need to have an array with order of rules
	setOfRules[3] = func(currentOutput []string) []string {
		result := append(currentOutput, "Fizz")
		return result
	}
	orderOfRules = append(orderOfRules, 3)

	setOfRules[5] = func(currentOutput []string) []string {
		result := append(currentOutput, "Buzz")
		return result
	}
	orderOfRules = append(orderOfRules, 5)

	for _, divisor := range orderOfRules {
		if number%divisor == 0 {
			effect := setOfRules[divisor]
			output = effect(output)
		}
	}

	return output
}

func formatOutput(output []string) string {
	return strings.Join(output[:], "")
}

func main() {
	for i := 0; i < 100; i++ {
		outputWords := determineOutput(i)
		output := formatOutput(outputWords)
		if len(output) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
	}

}
