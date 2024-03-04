package propertybased

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Digit uint16
	Roman string
}{
	{Digit: 1, Roman: "I"},
	{Digit: 2, Roman: "II"},
	{Digit: 3, Roman: "III"},
	{Digit: 4, Roman: "IV"},
	{Digit: 5, Roman: "V"},
	{Digit: 6, Roman: "VI"},
	{Digit: 7, Roman: "VII"},
	{Digit: 8, Roman: "VIII"},
	{Digit: 9, Roman: "IX"},
	{Digit: 10, Roman: "X"},
	{Digit: 14, Roman: "XIV"},
	{Digit: 18, Roman: "XVIII"},
	{Digit: 20, Roman: "XX"},
	{Digit: 39, Roman: "XXXIX"},
	{Digit: 40, Roman: "XL"},
	{Digit: 47, Roman: "XLVII"},
	{Digit: 49, Roman: "XLIX"},
	{Digit: 50, Roman: "L"},
	{Digit: 90, Roman: "XC"},
	{Digit: 100, Roman: "C"},
	{Digit: 400, Roman: "CD"},
	{Digit: 500, Roman: "D"},
	{Digit: 900, Roman: "CM"},
	{Digit: 1000, Roman: "M"},
	{Digit: 1984, Roman: "MCMLXXXIV"},
	{Digit: 3999, Roman: "MMMCMXCIX"},
	{Digit: 2014, Roman: "MMXIV"},
	{Digit: 1006, Roman: "MVI"},
	{Digit: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run("test digit to roman", func(t *testing.T) {
			got := ConvertToRoman(test.Digit)
			want := test.Roman

			assertEqualsTo(t, got, want)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Digit), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Digit

			assertEqualsTo(t, got, want)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(digit uint16) bool {
		if digit > 3999 { // due to implementation 3999 is the max val
			return true
		}
		t.Log("testing", digit)
		roman := ConvertToRoman(digit)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == digit
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func assertEqualsTo(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
