package numerals

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumeralsDesc = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var roman strings.Builder

	for _, numeral := range allRomanNumeralsDesc {
		for arabic >= numeral.Value {
			roman.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		roman.WriteString("X")
	// 		arabic -= 10
	// 	case arabic == 9:
	// 		roman.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		roman.WriteString("V")
	// 		arabic -= 5
	// 	case arabic == 4:
	// 		roman.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		roman.WriteString("I")
	// 		arabic -= 1
	// 	}
	// }

	return roman.String()
}
