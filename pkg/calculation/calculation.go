package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// Операции и их приоритет
var precedence = map[rune]int{
	'(': 0,
	')': 0,
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func isLeftParenthesis(r rune) bool {
	return r == '('
}

func isRightParenthesis(r rune) bool {
	return r == ')'
}

// Функция для преобразования инфиксной записи в постфиксную
func infixToPostfix(infix []rune) ([]string, error) {
	var output []string // Выходное выражение в постфиксной форме
	var stack []rune    // Стек для хранения операторов
	openBracketsCount := 0
	closeBracketsCount := 0

	for _, r := range infix {
		if isDigit(r) { // Если символ является цифрой, добавляем её к выходному списку
			output = append(output, string(r))
		} else if isOperator(r) { // Если оператор
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[r] {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				output = append(output, string(popped))
			}
			stack = append(stack, r)
		} else if isLeftParenthesis(r) {
			stack = append(stack, r)
			openBracketsCount++
		} else if isRightParenthesis(r) {
			for !isLeftParenthesis(stack[len(stack)-1]) {
				popped := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				output = append(output, string(popped))
				if len(stack) == 0 {
					return nil, errors.New("несбалансированные скобки")
				}
			}

			stack = stack[:len(stack)-1]
			closeBracketsCount++
		} else {
			return nil, fmt.Errorf("неизвестный символ %c", r)
		}
	}

	// Добавляем оставшиеся операторы из стека в выходной список
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		output = append(output, string(popped))
	}

	// Проверяем баланс скобок
	if openBracketsCount != closeBracketsCount {
		return nil, errors.New("несоответствие скобок: количество открывающих и закрывающих скобок различается")
	}

	return output, nil
}

// Вычисление значения выражения в постфиксной форме
func calculatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if isDigit([]rune(token)[0]) { // Если токен - число
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, value)
		} else { // Если токен - оператор
			if len(stack) < 2 {
				return 0, fmt.Errorf("ошибка в выражении")
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, fmt.Errorf("деление на ноль")
				}
				stack = append(stack, a/b)
			default:
				return 0, fmt.Errorf("неизвестный оператор %s", token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("ошибка в выражении")
	}

	return stack[0], nil
}

// Основная функция для вычисления выражения
func Calc(expression string) (string, error) {
	// Удаляем пробелы
	expression = regexp.MustCompile(`\s+`).ReplaceAllString(expression, "")

	infix := []rune(expression)
	postfix, err := infixToPostfix(infix)
	if err != nil {
		return "0", err
	}
	result, err := calculatePostfix(postfix)
	if err != nil {
		return "0", err
	}

	formattedResult := fmt.Sprintf("%.6g", math.Round(result*1e6)/1e6)

	return formattedResult, nil
}

func main() {
	expression := "2+2*(4- 2)"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(result)
	}
}
