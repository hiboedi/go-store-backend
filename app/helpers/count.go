package helpers

const (
	TaxPercent uint32  = 10 //10%
	discount   uint32  = 20
	price      float64 = 1000
	quantity   uint32  = 4
)

func CalculateDiscount(basePrice float64, discount uint32) float64 {
	discountRate := int(discount) / 100
	return basePrice * float64(discountRate)
}

func BasePrice(price float64, quantity uint32) float64 {
	return price * float64(quantity)
}

func PriceAfterDiscount(basePrice float64, discountAmount float64) float64 {
	return basePrice - discountAmount
}
func TaxAmount(priceAfterDiscount float64) float64 {
	return priceAfterDiscount * float64(TaxPercent) / 100
}

func TotalPrice(basePrice, discountAmount, taxAmount float64) float64 {
	return basePrice - discountAmount - taxAmount
}
