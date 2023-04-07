package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnit := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return grossUnit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; ok {
		bill[item] += units[unit]
		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Given unit is not in the units map
	if _, ok := units[unit]; !ok {
		return false
	}

	// Given item is not in the bill map
	if _, ok := bill[item]; !ok {
		return false
	}

	if bill[item] >= units[unit] {
		if bill[item] == units[unit] {
			delete(bill, item)
		} else {
			bill[item] -= units[unit]
		}
		return true
	}

	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if amount, ok := bill[item]; ok {
		return amount, ok
	}
	return 0, false
}
