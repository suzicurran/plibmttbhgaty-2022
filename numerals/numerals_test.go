package numerals

import "testing"

var cases = []struct {
	Arabic int
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 26, Roman: "XXVI"},
	{Arabic: 29, Roman: "XXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
}

func TestNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run("ConvertToArabic returns the arabic equivalent of a roman numeral", func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Arabic
			assertCorrectMessageInts(t, got, want)
		})

		t.Run("ConvertToRoman returns the roman equivalent of an arabic numeral", func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Roman
			assertCorrectMessageStrings(t, got, want)
		})
	}
}

func assertCorrectMessageStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but expected %q", got, want)
	}
}

func assertCorrectMessageInts(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d but expected %d", got, want)
	}
}
