package string_unpacker

import (
	"strings"
	"unicode"
)

const (
	nonDigitSymbol = iota
	anySymbol
	slashedSymbol
)

func Unpack(packedStr string) string {

	state := nonDigitSymbol

	previousSymbol := ""
	result := ""

	for _, c := range packedStr {
		switch state {
		case nonDigitSymbol:
			if unicode.IsDigit(c) {
				return ""
			}
			state = anySymbol
			fallthrough
		case anySymbol:
			if unicode.IsDigit(c) {
				result += strings.Repeat(previousSymbol, int(c-'0'))
				previousSymbol = ""
				state = nonDigitSymbol
			} else if c == '\\' {
				state = slashedSymbol
				result += previousSymbol
				previousSymbol = ""
			} else {
				result += previousSymbol
				previousSymbol = string(c)
			}
		case slashedSymbol:
			previousSymbol = string(c)
			state = anySymbol
		}
	}
	result += previousSymbol
	return result
}
