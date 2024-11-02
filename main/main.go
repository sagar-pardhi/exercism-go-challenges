package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Replace(log string, oldRune, newRune rune) string {
	oldRuneValue := fmt.Sprintf("%c", oldRune)
	newRuneValue := fmt.Sprintf("%c", newRune)

	return strings.Replace(log, oldRuneValue, newRuneValue, 1)
}

func main() {
	log := "please replace '👎' with '👍'"
	res := Replace(log, '👎', '👍')

	fmt.Println(res)

	test := "hello❗"
	fmt.Println(utf8.RuneCountInString(test) <= 6)
}
