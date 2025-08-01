package order

import "github.com/Js41313/Futuer-2/internal/types"

func getDiscount(discounts []types.SubscribeDiscount, inputMonths int64) float64 {
	var finalDiscount int64 = 100

	for _, discount := range discounts {
		if inputMonths >= discount.Quantity && discount.Discount < finalDiscount {
			finalDiscount = discount.Discount
		}
	}
	return float64(finalDiscount) / float64(100)
}
