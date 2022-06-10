package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	Amount, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += Amount
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	OldItemAmount, ItemExist := bill[item]
	NewItemAmount, UnitExist := units[unit]
	if !ItemExist || !UnitExist || NewItemAmount > OldItemAmount {
		return false
	}
	if NewItemAmount == OldItemAmount {
		delete(bill, item)
	} else {
		bill[item] -= NewItemAmount
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	ItemAmount, ItemExist := bill[item]
	if !ItemExist {
		return 0, false
	}
	return ItemAmount, true
}
