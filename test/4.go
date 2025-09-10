package main

// maxProfit finds the maximum profit from buying and selling stock.
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0] // Track min price seen
	maxProf := 0          // Track max profit
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price // Update min
		} else {
			prof := price - minPrice
			if prof > maxProf {
				maxProf = prof // Update max profit
			}
		}
	}
	return maxProf
}
