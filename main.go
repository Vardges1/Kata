package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operator func(int, int) int

var operators = map[string]operator{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"/": func(a, b int) int { return a / b },
	"*": func(a, b int) int { return a * b },
}

func parseArabicExpression(input string) (int, int, operator) {
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		panic("Invalid input format")
	}
	op, ok := operators[tokens[1]]
	if !ok {
		panic("Unknown operator: " + tokens[1])
	}
	a, err := strconv.Atoi(tokens[0])
	if err != nil || a < 1 || a > 10 {
		panic("Invalid operand: " + tokens[0])
	}
	b, err := strconv.Atoi(tokens[2])
	if err != nil || b < 1 || b > 10 {
		panic("Invalid operand: " + tokens[2])
	}
	return a, b, op
}

func intToRoman(n int) string {
	if n < 1 || n > 3999 {
		panic("Invalid Roman numeral")
	}
	roman := []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var result strings.Builder
	for _, r := range roman {
		for n >= r.value {
			result.WriteString(r.symbol)
			n -= r.value
		}
	}
	return result.String()
}

func parseRomanExpression(input string) (int, int, operator) {
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		panic("Недопустимым форматом ввода:")
	}
	op, ok := operators[tokens[1]]
	if !ok {
		panic("Ошибка: " + tokens[1])
	}
	a := romanToInt(tokens[0])
	if a < 1 || a > 10 {
		panic("Недопустимое значение: " + tokens[0])
	}
	b := romanToInt(tokens[2])
	if b < 1 || b > 10 {
		panic("Недопустимое значение: " + tokens[2])
	}
	return a, b, op
}

func romanToInt(roman string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	var prevValue int
	for _, r := range roman {
		value, ok := romanValues[r]
		if !ok {
			panic("Ошибка!")
		}
		if prevValue < value {
			result -= prevValue
			value -= prevValue
		}
		result += value
		prevValue = value
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите пример: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a, b int
	var op operator
	var isRoman bool
	if strings.ContainsAny(input, "IVXLCMD") {
		isRoman = true
	}
	if isRoman {
		a, b, op = parseRomanExpression(input)
	} else {
		a, b, op = parseArabicExpression(input)
	}
	result := op(a, b)
	if isRoman {
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}
}
//Пример ввода:5 + 6
// Вывод программы :11
