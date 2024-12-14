package hw02unpackstring

import (
	"errors"
	"slices"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	r := []rune(s)
	res, err := rebuildString(r)
	return res, err
}

func rebuildString(r []rune) (string, error) {
	numbers := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	result := make([]rune, 0)
	for i, v := range r {
		if i == 0 && slices.Contains(numbers, v) {
			return "", ErrInvalidString
		}

		if slices.Contains(numbers, r[i]) && slices.Contains(numbers, r[i-1]) {
			return "", ErrInvalidString
		}

		if !slices.Contains(numbers, v) {
			result = append(result, r[i])
		}

		if i != 0 && slices.Contains(numbers[1:], v) {
			l, err := strconv.Atoi(string(v))
			if err != nil {
				return "", ErrInvalidString
			}
			for j := 0; j <= l-2; j++ {
				result = append(result, r[i-1])
			}
		}

		if i != 0 && v == numbers[0] {
			result = result[:len(result)-1]
		}
	}
	return string(result), nil
}
