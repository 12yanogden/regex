package regex

import (
	"regexp"

	ibytes "github.com/12yanogden/bytes"
)

func Pretty(regex string) string {
	regexBytes := []byte(regex)
	pretty := []byte{}
	tabLevel := 0
	var handled byte

	for _, symbol := range regexBytes {
		switch symbol {
		case '(':
			if tabLevel > 0 && handled == '(' {
				pretty = append(pretty, '\n')
				pretty = ibytes.AppendRepeat(pretty, '\t', tabLevel)
			}

			pretty = append(pretty, symbol)

			tabLevel++

		case ')':
			tabLevel--
			pretty = append(pretty, '\n')
			pretty = ibytes.AppendRepeat(pretty, '\t', tabLevel)
			pretty = append(pretty, symbol)

		default:
			if handled == '(' {
				pretty = append(pretty, '\n')
				pretty = ibytes.AppendRepeat(pretty, '\t', tabLevel)
			}

			pretty = append(pretty, symbol)
		}

		handled = symbol
	}

	return string(pretty)
}

func Ugly(pretty string) string {
	pretty = regexp.MustCompile("[\t ]*//.*").ReplaceAllString(pretty, "")
	pretty = regexp.MustCompile("\n").ReplaceAllString(pretty, "")
	pretty = regexp.MustCompile("[\t]*"+`\(`).ReplaceAllString(pretty, "(")
	pretty = regexp.MustCompile(`\(`+"[\t]*").ReplaceAllString(pretty, "(")
	pretty = regexp.MustCompile("[\t]*"+`\)`).ReplaceAllString(pretty, ")")
	pretty = regexp.MustCompile(`\)`+"[\t]*").ReplaceAllString(pretty, ")")

	return pretty
}
