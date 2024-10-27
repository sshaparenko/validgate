package service

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

var ErrLuhnCheck error = errors.New("card number fails Luhn check")
var ErrExpDate error = errors.New("card is expired")

func VailedateNumber(number string) (bool, error) {
	return LuhnCheck(number)
}

// LuhnCheck validates card number according to Luhn algorithm
func LuhnCheck(number string) (bool, error) {
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

	return false, ErrLuhnCheck
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

func CheckExpDate(month int, year int) bool {
	currentYear, currentMonth, _ := time.Now().Date()

	if year > currentYear {
		return true
	}

	if year == currentYear && month > int(currentMonth) {
		return true
	}

	return false
}
