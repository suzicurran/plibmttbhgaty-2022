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

var romanNumeralPairs = []RomanNumeral{
	{900, "CM"},
	{400, "CD"},
	{90, "XC"},
	{40, "XL"},
	{9, "IX"},
	{4, "IV"},
}

var romanNumeralSingletons = []RomanNumeral{
	{1000, "M"},
	{500, "D"},
	{100, "C"},
	{50, "L"},
	{10, "X"},
	{5, "V"},
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

	return roman.String()
}

func ConvertToArabic(roman string) int {
	var arabic int = 0

	for _, numeralPair := range romanNumeralPairs {
		if strings.Index(roman, numeralPair.Symbol) != -1 {
			roman = strings.Replace(roman, numeralPair.Symbol, "", 1)
			arabic += numeralPair.Value
		}
	}

	for _, numeral := range romanNumeralSingletons {
		for strings.Index(roman, numeral.Symbol) != -1 {
			roman = strings.Replace(roman, numeral.Symbol, "", 1)
			arabic += numeral.Value
		}
	}

	return arabic
}
