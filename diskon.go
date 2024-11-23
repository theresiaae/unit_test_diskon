package main

func CalculateDiscount(originalPrice, discountPercent float64) float64 {
	if discountPercent < 0 || discountPercent > 100 {
		return -1
	}
	discount := originalPrice * discountPercent / 100
	return originalPrice - discount
}
