package main

func main() {

}

func romanNumerals(n int) string {

	if n == 5 {
		return "V"
	}

	if n == 4 {
		return "IV"
	}

	if n == 9 {
		return "IX"
	}

	roman := ""

	fives := n / 5
	for i := 0; i < fives; i++ {
		roman += "V"
	}

	for i := 0; i < (n - (fives * 5)); i++ {
		roman += "I"
	}

	return roman

}
