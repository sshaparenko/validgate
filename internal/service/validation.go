package service

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func VailedateNumber(number string) (bool, error) {
	return luhnCheck(number)
}

func luhnCheck(number string) (bool, error) {
	var sum int

	number = strings.Replace(number, " ", "", -1)
	if len(number) == 0 {
		return false, errors.New("card number is empty")
	}

	last_digit, err := getLastDigit(number)
	if err != nil {
		return false, fmt.Errorf("in service.ValidateNumber: %w", err)
	}

	cardDigits, err := convertStringToInts(number)
	if err != nil {
		return false, err
	}
	cardDigits = cardDigits[:len(number)-1]
	slices.Reverse(cardDigits)
	multiplyOddDigits(cardDigits)
	substractByNine(cardDigits)

	for _, n := range cardDigits {
		sum += n
	}

	mod := sum % 10

	if 10-mod == last_digit {
		return true, nil
	}

	return false, errors.New("card number fails Luhn check")
}

// convertSrtingToIntes converts card number
// of type string to an array of integers
func convertStringToInts(number string) ([]int, error) {
	card_len := len(number)
	cardDigits := make([]int, card_len)

	for i, n := range number {
		digit, err := strconv.Atoi(string(n))
		if err != nil {
			return nil, fmt.Errorf("in convertStringtoInts: %w", err)
		}
		cardDigits[i] = digit
	}

	return cardDigits, nil
}

func multiplyOddDigits(cardDigits []int) {
	for i := 0; i < len(cardDigits); i = i + 2 {
		cardDigits[i] = cardDigits[i] * 2
	}
}

func substractByNine(cardDigits []int) {
	for i, n := range cardDigits {
		if n > 9 {
			cardDigits[i] = n - 9
		}
	}
}

func getLastDigit(number string) (int, error) {
	lastRune := number[len(number)-1]
	return strconv.Atoi(string(lastRune))
}
