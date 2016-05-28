package klondike

func Move(from, to interface{}, index int) error {
	switch from := from.(type) {
	case *Stock:
		switch to := to.(type) {
		case *Waste:
			return StockToWaste(from, to, index)
		}
	case *Waste:
		switch to := to.(type) {
		case *Stock:
			return WasteToStock(from, to)
		}
	}
	return InvalidArgs
}

// 山札から捨て札
func StockToWaste(from *Stock, to *Waste, index int) error {
	if index >= len(*from) {
		return InvalidArgs
	}
	*to = append(*to, (*from)[index:]...)
	*from = (*from)[:index]
	return nil
}

// 捨て札を山札に
func WasteToStock(from *Waste, to *Stock) error {
	*to = append(*to, (*from)[0:]...)
	*from = (*from)[:0]
	return nil
}

func WasteToColumn(from *Waste, to *Column) error {
	return nil
}
