package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}
	if balance < 1000 {
		return 0.5
	}
	if balance < 5000 {
		return 1.621
	}

	// More 5k
	return 2.475
	// panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	ir := InterestRate(balance)
	balance = balance * float64(ir) / 100
	return balance
	// panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	balance += Interest(balance)
	return balance
	// panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var reqYears int
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		reqYears++
	}
	return reqYears
	// panic("Please implement the YearsBeforeDesiredBalance function")
}
