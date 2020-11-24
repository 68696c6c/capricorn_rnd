package utils

import (
	"strings"

	"github.com/jinzhu/inflection"
)

var specialChars = []rune{
	'@',
	'#',
	'$',
	'%',
	'^',
	'&',
	'*',
	'(',
	')',
	':',
	';',
	'<',
	'>',
	',',
	'?',
	'"',
	'\'',
	'+',
	'=',
	'~',
	'`',
}

func Singular(input string) string {
	return inflection.Singular(input)
}

func Plural(input string) string {
	return inflection.Plural(input)
}

func Kebob(separatedInput ...string) string {
	return separatedToKebob(separatedInput...)
}

func Snake(separatedInput ...string) string {
	return separatedToSnake(separatedInput...)
}

func Space(separatedInput ...string) string {
	return separatedToSpace(separatedInput...)
}

func Pascal(separatedInput ...string) string {
	return separatedToPascal(separatedInput...)
}

func Camel(separatedInput ...string) string {
	return separatedToCamel(separatedInput...)
}

func isSpecialChar(r rune) bool {
	for _, c := range specialChars {
		if c == r {
			return true
		}
	}
	return false
}

func clean(input ...string) string {
	separator := '-'
	characters := strings.Join(input, string(separator))
	var result strings.Builder
	for i := 0; i < len(characters); i++ {
		b := characters[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' || b == '-' || b == '_' {
			result.WriteByte(b)
		} else if isSpecialChar(rune(b)) {
			result.WriteRune(separator)
		}
	}
	return result.String()
}

func separatedToCamel(input ...string) string {
	return separatedToMedial(clean(input...), false)
}

func separatedToPascal(input ...string) string {
	return separatedToMedial(clean(input...), true)
}

func separatedToSpace(input ...string) string {
	return separatedToSeparated(clean(input...), ' ')
}

func separatedToSnake(input ...string) string {
	return separatedToSeparated(clean(input...), '_')
}

func separatedToKebob(input ...string) string {
	return separatedToSeparated(clean(input...), '-')
}

func separatedToMedial(input string, leadingCap bool) string {
	isToUpper := false
	var output string
	for k, v := range input {
		if k == 0 && leadingCap {
			output = strings.ToUpper(string(input[0]))
		} else {
			if isToUpper {
				output += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' || v == '-' || v == ' ' {
					isToUpper = true
				} else {
					output += string(v)
				}
			}
		}
	}
	return output
}

func separatedToSeparated(input string, separator rune) string {
	var output string
	for _, v := range input {
		if v == '_' || v == '-' || v == ' ' {
			output += string(separator)
		} else {
			output += strings.ToLower(string(v))
		}
	}
	return output
}
