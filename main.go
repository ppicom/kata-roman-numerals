package main

func main() {

}

func romanNumerals(n int) string {
	if n == 4 {
		return "IV"
	}

	if n == 9 {
		return "IX"
	}

	roman := ""
	remainder := n

	fifties := n / 50
	for i := 0; i < fifties; i++ {
		roman += "L"
	}
	remainder -= fifties * 50

	if remainder >= 40 && remainder%40 < 10 {
		roman = "XL" + roman
		remainder -= 40
	}

	tens := remainder / 10
	for i := 0; i < tens; i++ {
		roman += "X"
	}
	remainder -= tens * 10

	if remainder == 4 {
		roman += "IV"
		return roman
	}

	if remainder == 9 {
		roman += "IX"
		return roman
	}

	fives := remainder / 5
	for i := 0; i < fives; i++ {
		roman += "V"
	}
	remainder -= fives * 5

	for i := 0; i < remainder; i++ {
		roman += "I"
	}

	return roman

}
