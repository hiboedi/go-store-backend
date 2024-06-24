package helpers

var TaxPercent = 10
var DiscountPercent = 15

func GetDiscountAmount(price float64) float64 {
	return price / float64(DiscountPercent)
}

func GetTaxAmount(price float64) float64 {
	return price / float64(TaxPercent)
}
