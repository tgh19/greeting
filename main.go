package main

import (
	"strings"
	"unicode"
)

func Greet(names []string) string {
	var name interface{}
	var endsInExclamation bool
	lenNames := len(names)

	// Requirement 4
	if lenNames == 1 {
		name = names[0]
	}

	// Requirement 4
	if lenNames >= 2 {
		var newNames []string
		for _, name := range names {
			var splitNames []string

			// Requirement 8
			if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
				newNames = append(newNames, strings.Replace(name, "\"", "", -1))

				// Requirement 7
			} else if strings.Contains(name, ",") {
				splitNames = strings.Split(name, ",")
				for _, splitName := range splitNames {
					newNames = append(newNames, splitName)
				}
			} else {
				newNames = append(newNames, name)
			}
		}
		names = newNames

		var upperNames []string
		var lowerNames []string
		for _, name := range names {
			if IsUpper(name) {
				upperNames = append(upperNames, name)
			} else {
				lowerNames = append(lowerNames, name)
			}
		}

		// Requirement 5, 6
		lenLowerNames := len(lowerNames)
		var combinedLowerNames string
		if lenLowerNames == 1 {
			combinedLowerNames = lowerNames[0]
		} else if lenLowerNames == 2 {
			combinedLowerNames = lowerNames[0] + " and " + lowerNames[1]
			name = combinedLowerNames
		} else if lenLowerNames >= 3 {
			for i, name := range lowerNames {
				if i == lenLowerNames-2 {
					combinedLowerNames += name + ", and "
				} else if i == lenLowerNames-1 {
					combinedLowerNames += name
				} else {
					combinedLowerNames += name + ", "
				}
			}
		}
		name = combinedLowerNames

		// Requirement 5, 6
		var combinedUpperNames string
		lenUpperNames := len(upperNames)
		if lenUpperNames == 1 {
			combinedUpperNames = upperNames[0]
		} else if lenUpperNames == 2 {
			combinedUpperNames = upperNames[0] + " AND " + upperNames[1]
			name = combinedUpperNames
		} else if lenUpperNames >= 3 {
			for i, name := range upperNames {
				if i == lenUpperNames-2 {
					combinedUpperNames += name + ", AND "
				} else if i == lenUpperNames-1 {
					combinedUpperNames += name
				} else {
					combinedUpperNames += name + ", "
				}
			}
		}
		// Determine how to combine the previous output if it exists
		if lenLowerNames > 0 && lenUpperNames > 0 {
			name = combinedLowerNames + ". AND HELLO " + combinedUpperNames
			endsInExclamation = true
		} else if lenUpperNames > 0 {
			name = "HELLO " + combinedUpperNames
			endsInExclamation = true
		}
	}

	// Requirement 2
	if name == nil {
		name = "my friend"
	}

	// Requirement 3
	if IsUpper(name.(string)) {
		return "HELLO " + name.(string) + "!"
	}

	// Needed for Requirement 5, 6
	endChar := "."
	if endsInExclamation {
		endChar = "!"
	}

	// Requirement 1
	return "Hello, " + name.(string) + endChar
}

func IsUpper(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) && !unicode.IsUpper(char) {
			return false
		} else if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
