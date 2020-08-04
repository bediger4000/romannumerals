package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	romanstringrep := os.Args[1]
	n, err := romanToDecimal(romanstringrep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %d\n", romanstringrep, n)
}

func romanToDecimal(romanstringrep string) (int, error) {
	numerals := []rune(romanstringrep)

	lastValue := 0
	sum := 0
	run := 0
	for i := range numerals {
		numeral := unicode.ToUpper(numerals[i])
		var value int
		switch numeral {
		case 'M':
			value = 1000
		case 'D':
			value = 500
		case 'C':
			value = 100
		case 'L':
			value = 50
		case 'X':
			value = 10
		case 'V':
			value = 5
		case 'I':
			value = 1
		default:
			return 0, fmt.Errorf("found non-digit %c", numerals[i])
		}

		if lastValue == 0 {
			run = value
		} else {
			if value == lastValue {
				run += value
			} else if value < lastValue {
				value += run
				run = value
			} else {
				// value > lastValue
				sum += (value - run)
				run = 0
			}
		}

		lastValue = value
	}

	if run != 0 {
		sum += run
	}

	return sum, nil
}
