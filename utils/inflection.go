package utils

import "strings"

type Inflection struct {
	Kebob  string
	Snake  string
	Space  string
	Pascal string
	Camel  string
}

func NewInflection(separated string) Inflection {
	return Inflection{
		Kebob:  Kebob(separated),
		Snake:  Snake(separated),
		Space:  Space(separated),
		Pascal: Pascal(separated),
		Camel:  Camel(separated),
	}
}

func Kebob(separated string) string {
	return separatedToKebob(separated)
}

func Snake(separated string) string {
	return separatedToSnake(separated)
}

func Space(separated string) string {
	return separatedToSpace(separated)
}

func Pascal(separated string) string {
	return separatedToPascal(separated)
}

func Camel(separated string) string {
	return separatedToCamel(separated)
}

func separatedToCamel(input string) string {
	return separatedToMedial(input, false)
}

func separatedToPascal(input string) string {
	return separatedToMedial(input, true)
}

func separatedToSpace(input string) string {
	return separatedToSeparated(input, ' ')
}

func separatedToSnake(input string) string {
	return separatedToSeparated(input, '_')
}

func separatedToKebob(input string) string {
	return separatedToSeparated(input, '-')
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
				if v == '_' || v == '-' {
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
		if v == '_' || v == '-' {
			output += string(separator)
		} else {
			output += strings.ToLower(string(v))
		}
	}
	return output
}
