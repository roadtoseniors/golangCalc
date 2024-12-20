package calc

import (
	"errors"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	index := 0
	result, err := parseExpression(expression, &index)
	if err != nil {

		return 0, errors.New("Expression is not valid")
	}

	for index < len(expression) {
		if expression[index] != ' ' {
			return 0, errors.New("Expression is not valid")
		}
		index++
	}

	return result, nil
}

func parseExpression(expression string, index *int) (float64, error) {
	result, err := parseTerm(expression, index)
	if err != nil {
		return 0, err
	}

	for *index < len(expression) {
		char := expression[*index]

		if char == ' ' {
			*index++
			continue
		}

		if char == '+' || char == '-' {
			*index++
			nextTerm, err := parseTerm(expression, index)
			if err != nil {
				return 0, err
			}
			if char == '+' {
				result += nextTerm
			} else {
				result -= nextTerm
			}
		} else {
			break
		}
	}

	return result, nil
}

func parseTerm(expression string, index *int) (float64, error) {
	result, err := parseFactor(expression, index)
	if err != nil {
		return 0, err
	}

	for *index < len(expression) {
		char := expression[*index]

		if char == ' ' {
			*index++
			continue
		}

		if char == '*' || char == '/' {
			*index++
			nextFactor, err := parseFactor(expression, index)
			if err != nil {
				return 0, err
			}
			if char == '*' {
				result *= nextFactor
			} else {
				if nextFactor == 0 {
					return 0, errors.New("деление на ноль")
				}
				result /= nextFactor
			}
		} else {
			break
		}
	}

	return result, nil
}

func parseFactor(expression string, index *int) (float64, error) {
	if *index < len(expression) && expression[*index] == '(' {
		*index++
		result, err := parseExpression(expression, index)
		if err != nil {
			return 0, err
		}
		if *index >= len(expression) || expression[*index] != ')' {
			return 0, errors.New("отсутствует закрывающая скобка")
		}
		*index++ // Пропускаем ')'
		return result, nil
	}

	return readNumber(expression, index) // Если это число, читаем его
}

// readNumber считывает число из строки.
func readNumber(expression string, index *int) (float64, error) {
	start := *index
	for *index < len(expression) && (unicode.IsDigit(rune(expression[*index])) || expression[*index] == '.') {
		*index++
	}
	if start == *index {
		return 0, errors.New("ожидалось число")
	}
	return strconv.ParseFloat(expression[start:*index], 64)
}
