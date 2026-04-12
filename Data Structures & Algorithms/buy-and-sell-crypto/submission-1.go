func maxProfit(prices []int) int {
lowPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < lowPrice {
			lowPrice = price
		} else {
			if price-lowPrice > maxProfit {
				maxProfit = price - lowPrice
			}
		}
	}

	return maxProfit
}
