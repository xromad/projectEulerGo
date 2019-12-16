package num2words

import (
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{0, "zero"},
		{1, "one"},
		{12, "twelve"},
		{123, "one hundred twenty three"},
		{1230, "one thousand two hundred thirty"},
		{1234, "one thousand two hundred thirty four"},
		// {12345, "twelve thousand three hundred fourty five"},
		// {123456, "one hundred twenty three thousand four hundred fifty six"},
		// {1234567, "one million two hundred thirty four thousand five hundred sixty seven"},
		// {12345678, "twelve million three hundred fourty five thousand six hundred seventy eight"},
		// {123456789, "one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine"},
		// {987654321, "nine hundred eighty seven million six hundred fifty four thousand three hundred twenty one"
		{10000, ""},
	}

	for _, c := range testCases {
		got := Convert(c.num)
		if got != c.want {
			t.Errorf("Convert(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}

func TestOnes(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{0, "zero"},
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
	}

	for _, c := range testCases {
		got := ones(c.num)
		if got != c.want {
			t.Errorf("ones(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}

func TestTens(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{10, "ten"},
		{11, "eleven"},
		{12, "twelve"},
		{13, "thirteen"},
		{14, "fourteen"},
		{15, "fifteen"},
		{16, "sixteen"},
		{17, "seventeen"},
		{18, "eighteen"},
		{19, "nineteen"},
		{20, "twenty"},
		{21, "twenty one"},
		{30, "thirty"},
		{32, "thirty two"},
		{40, "fourty"},
		{43, "fourty three"},
		{50, "fifty"},
		{54, "fifty four"},
		{60, "sixty"},
		{65, "sixty five"},
		{70, "seventy"},
		{76, "seventy six"},
		{80, "eighty"},
		{87, "eighty seven"},
		{90, "ninety"},
		{98, "ninety eight"},
	}

	for _, c := range testCases {
		got := tens(c.num)
		if got != c.want {
			t.Errorf("tens(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}

func TestHundreds(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{100, "one hundred"},
		{101, "one hundred one"},
		{112, "one hundred twelve"},
		{123, "one hundred twenty three"},
		{200, "two hundred"},
		{300, "three hundred"},
		{400, "four hundred"},
		{500, "five hundred"},
		{600, "six hundred"},
		{700, "seven hundred"},
		{800, "eight hundred"},
		{900, "nine hundred"},
		{999, "nine hundred ninety nine"},
		{1000, ""},
	}

	for _, c := range testCases {
		got := hundreds(c.num)
		if got != c.want {
			t.Errorf("hundreds(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}

func TestThousands(t *testing.T) {
	testCases := []struct {
		num  int
		want string
	}{
		{1000, "one thousand"},
		{1001, "one thousand one"},
		{1012, "one thousand twelve"},
		{1023, "one thousand twenty three"},
		{1123, "one thousand one hundred twenty three"},
		{2000, "two thousand"},
		{3000, "three thousand"},
		{4000, "four thousand"},
		{5000, "five thousand"},
		{6000, "six thousand"},
		{7000, "seven thousand"},
		{8000, "eight thousand"},
		{9000, "nine thousand"},
		{9999, "nine thousand nine hundred ninety nine"},
		{10000, ""},
	}

	for _, c := range testCases {
		got := thousands(c.num)
		if got != c.want {
			t.Errorf("thousands(%v) == %s, want %s", c.num, got, c.want)
		}
	}
}
