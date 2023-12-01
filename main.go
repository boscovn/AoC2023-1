package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	keys := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	wordToNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		text := scanner.Text()
		num := 0
		if text == "" {
			break
		}
		text = text + "y"
		runes := []rune(text)
	outerLoop:
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				num = int(runes[i] - '0')
				break
			}
			for _, key := range keys {
				if i-len(key) >= 0 {
					if string(runes[i-len(key):i]) == key {
						num = wordToNumber[key]
						break outerLoop
					}
				}
			}

		}
	otherOuterLoop:
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				num += (int(runes[i]-'0') * 10)
				break
			}
			for _, key := range keys {
				if (i + len(key)) <= len(runes) {
					if string(runes[i:i+len(key)]) == key {
						num += wordToNumber[key] * 10
						break otherOuterLoop
					}
				}
			}

		}
		sum += num

	}

	fmt.Println(sum)
}
