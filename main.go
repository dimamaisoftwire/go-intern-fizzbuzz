package main

import (
	"fmt"
	"strings"
)

var Rules []Rule

type Rule struct {
	Divisor int
	Effect  func([]string) []string
}

func determineOutput(number int) []string {

	Rules = make([]Rule, 0)
	output := make([]string, 0)

	threeRule := Rule{3, func(currentOutput []string) []string {
		result := append(currentOutput, "Fizz")
		return result
	}}

	fiveRule := Rule{5, func(currentOutput []string) []string {
		result := append(currentOutput, "Buzz")
		return result
	}}
	sevenRule := Rule{7, func(currentOutput []string) []string {
		result := append(currentOutput, "Bang")
		return result
	}}

	elevenRule := Rule{11, func(currentOutput []string) []string {
		result := []string{"Bong"}
		return result
	}}

	thirteenRule := Rule{13, func(currentOutput []string) []string {
		var result []string
		for i, word := range currentOutput {
			if word[0] == 'B' {
				result = append(currentOutput[:i], append([]string{"Fezz"}, currentOutput[i:]...)...)
				return result
			}
		}
		result = append(currentOutput, "Fezz")
		return result
	}}

	seventeenRule := Rule{17, func(currentOutput []string) []string {
		var result []string
		for i, j := 0, len(currentOutput)-1; i < j; i, j = i+1, j-1 {
			currentOutput[i], currentOutput[j] = currentOutput[j], currentOutput[i]
		}
		result = currentOutput
		return result
	}}
	Rules = append(Rules, threeRule, fiveRule, sevenRule, elevenRule, thirteenRule, seventeenRule)
	for _, rule := range Rules {
		if number%rule.Divisor == 0 {
			output = rule.Effect(output)
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
