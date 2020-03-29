package mycalculator

import (
	"fmt"
	"strconv"
	"strings"
)


type calc struct{}

func (calc) Operate(value1 int, value2 int, operator string) int {

	switch operator {
	case "+":
		fmt.Println(value1 + value2)
		return value1 + value2

	case "*":
		fmt.Println(value1 * value2)
		return value1 * value2
	case "-":
		fmt.Println(value1 - value2)
		return value1 - value2
	case "/":
		fmt.Println(value1 / value2)
		return value1 / value2
	default:
		fmt.Println("INVALID OPERATOR")
		return 0
	}
}

func parseInfo(valuesint []string, operator string) (int, int, string) {
	i1, _ := strconv.Atoi(valuesint[0])
	i2, _ := strconv.Atoi(valuesint[1])
	return i1, i2, operator
}

func ReadInput() string {
	fmt.Println("Please enter your operation:")
	var input string
	fmt.Scanln(&input)
	return input
}

func SplitInfo(entrada string) (int, int, string) {
	if strings.Contains(entrada, "+") {
		arrResult := strings.Split(entrada, "+")
		return parseInfo(arrResult, "+")
	}
	if strings.Contains(entrada, "*") {
		arrResult := strings.Split(entrada, "*")
		return parseInfo(arrResult, "*")
	}
	if strings.Contains(entrada, "-") {
		arrResult := strings.Split(entrada, "-")
		return parseInfo(arrResult, "-")
	}
	if strings.Contains(entrada, "/") {
		arrResult := strings.Split(entrada, "/")
		return parseInfo(arrResult, "/")
	}
	return 0, 0, "INVALID"
}
