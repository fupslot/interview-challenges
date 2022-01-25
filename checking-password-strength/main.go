package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isLower(r rune) bool {
	return unicode.IsLower(r)
}

func isUpper(r rune) bool {
	return unicode.IsUpper(r)
}

var special = "!@#$%^&*()_+{}[]<|>/~!"

func isSpecial(r rune) bool {
	return strings.ContainsRune(special, r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func any(v []int8) bool {
	for _, i := range v {
		if i > 0 {
			return true
		}
	}

	return false
}
func toBit(b bool) int8 {
	if b {
		return 1
	}

	return 0
}

func inCommonPasswordList(word string) bool {
	f, err := os.OpenFile("common_password_list.txt", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == word {
			return true
		}
	}

	return false
}

func main() {
	var word string = os.Args[1:][0]

	if inCommonPasswordList(word) {
		fmt.Printf("password [%s] in the common password list. Score(0)", word)
		os.Exit(0)
	}

	var score int8

	var lowercases []int8 = make([]int8, len(word))
	var uppercases []int8 = make([]int8, len(word))
	var specials []int8 = make([]int8, len(word))
	var digits []int8 = make([]int8, len(word))

	for i, r := range word {
		lowercases[i] = toBit(isLower(r))
		uppercases[i] = toBit(isUpper(r))
		specials[i] = toBit(isSpecial(r))
		digits[i] = toBit(isDigit(r))
	}

	if len(word) > 8 {
		score += 1
	}

	if len(word) > 15 {
		score += 1
	}

	if len(word) > 20 {
		score += 1
	}

	if any(lowercases) {
		score += 1
	}

	if any(specials) {
		score += 1
	}

	if any(uppercases) {
		score += 1
	}

	if any(digits) {
		score += 1
	}

	fmt.Printf("password: [%s] Score (%d/7)\n", word, score)
}
