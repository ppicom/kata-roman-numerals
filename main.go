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

	roman := ""
	for i := 0; i < n; i++ {
		roman += "I"
	}

	return roman

}
