package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var resultString strings.Builder

	if len(input) == 0 {
		return resultString.String(), nil
	}

	inputRune := []rune(input)

	// если первый символ цифра, то это ошибка
	if _, errCurrent := strconv.Atoi(string(inputRune[0])); errCurrent == nil {
		return "", ErrInvalidString
	}

	// если последний символ слэш, то это ошибка
	if inputRune[len(inputRune)-1] == 92 {
		return "", ErrInvalidString
	}

	var i int
	for i = 0; i < len(inputRune)-1; i++ {
		var simvol string
		multi := 1

		// проверяем, может быть следующий символ цифра
		nextRuneAsNumber, errNext := strconv.Atoi(string(inputRune[i+1]))

		if inputRune[i] == 92 {
			// экранировать можем только слэш или цифру
			if inputRune[i+1] != 92 && errNext != nil {
				return "", ErrInvalidString
			}
			// берем следующие символы
			simvol = string(inputRune[i+1])
			i++
			if i < len(inputRune)-1 {
				nextRuneAsNumber, errNext = strconv.Atoi(string(inputRune[i+1]))
			}
		} else {
			simvol = string(inputRune[i])
		}

		// следующий символ цифра
		if errNext == nil && i < len(inputRune)-1 {
			// но не двузначное число
			if i < len(inputRune)-2 {
				if _, errNextNext := strconv.Atoi(string(inputRune[i+2])); errNextNext == nil {
					return "", ErrInvalidString
				}
			}
			multi = nextRuneAsNumber
			i++
		}

		resultString.WriteString(strings.Repeat(simvol, multi))
	}

	// последний символ просто дописываем в итоговую строку
	if i == len(inputRune)-1 {
		resultString.WriteRune(inputRune[i])
		return resultString.String(), nil
	}

	return resultString.String(), nil
}
