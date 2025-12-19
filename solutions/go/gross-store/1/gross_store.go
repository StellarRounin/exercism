package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int{}

	measurements["quarter_of_a_dozen"] = 3
	measurements["half_of_a_dozen"] = 6
	measurements["dozen"] = 12
	measurements["small_gross"] = 120
	measurements["gross"] = 144
	measurements["great_gross"] = 1728

	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	_, itemExists := bill[item]

	if !unitExists {
		return false
	}

	if itemExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemExists := bill[item]
	unitValue, unitExists := units[unit]

	if !itemExists || !unitExists {
		return false
	}

	if itemValue-unitValue < 0 {
		return false
	} else if itemValue-unitValue == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= unitValue
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, itemExists := bill[item]

	if !itemExists {
		return 0, false
	}

	return bill[item], true
}
