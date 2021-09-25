func intToRoman(num int) string {
	romanNumerals := ""

	if num/1000 > 0 {
		for i := 1; i <= (num / 1000); i++ {
			romanNumerals += "M"
		}
		num = num % 1000	
	}

	if  num >= (1000-100) {
		romanNumerals += "CM"
		num = num % (1000-100)
	}

	if num/500 > 0 {
		for i := 1; i <= (num / 500); i++ {
			romanNumerals += "D"
		}
		num = num % 500
	}

	if num >= (500-100) {
		romanNumerals += "CD"
		num = num % (500-100)
	}

	if num/100 > 0 {
		for i := 1; i <= (num / 100); i++ {
			romanNumerals += "C"
		}
		num = num % 100
	}

	if num >= (100-10) {
		romanNumerals += "XC"
		num = num % (100-10)
	}

	if num/50 > 0 {
		for i := 1; i <= (num / 50); i++ {
			romanNumerals += "L"
		}
		num = num % 50
	}

	if num >= (50-10) {
		romanNumerals += "XL"
		num = num % (50-10)
	}

	if num/10 > 0 {
		for i := 1; i <= (num / 10); i++ {
			romanNumerals += "X"
		}
		num = num % 10
	}
	
	if num == (10-1) {
		romanNumerals += "IX"
		num = num % (10-1)
	}

	if num/5 > 0 {
		for i := 1; i <= (num / 5); i++ {
			romanNumerals += "V"
		}
		num = num % 5
	}

	if num == (5-1) {
		romanNumerals += "IV"
		num = num % (5-1)
	}

	for i := 1; i <= (num); i++ {
		romanNumerals += "I"
	}

	return romanNumerals
}