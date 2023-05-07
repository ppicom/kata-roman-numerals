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

	dees := n / 500
	for i := 0; i < dees; i++ {
		roman += "D"
	}
	remainder -= dees * 500

	if remainder >= 400 && remainder%400 < 100 {
		roman += "CD"
		remainder -= 400
	}

	centurions := remainder / 100
	for i := 0; i < centurions; i++ {
		roman += "C"
	}
	remainder -= centurions * 100

	if remainder >= 90 && remainder%90 < 10 {
		roman = "XC" + roman
		remainder -= 90
	}

	fifties := remainder / 50
	for i := 0; i < fifties; i++ {
		roman += "L"
	}
	remainder -= fifties * 50

	if remainder >= 40 && remainder%40 < 10 {
		roman += "XL"
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
