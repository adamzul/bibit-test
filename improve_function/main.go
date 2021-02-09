package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("adam(bibit)zulkarnain"))
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		indexClosingBracketFound := strings.Index(str, ")")
		if indexFirstBracketFound >= 0 && indexClosingBracketFound >= 0 && indexClosingBracketFound > indexFirstBracketFound {
			runes := []rune(str)
			return string(runes[indexFirstBracketFound+1 : indexClosingBracketFound])
		}
	}
	return ""
}
