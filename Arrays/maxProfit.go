package Arrays

import "math"

func maxProfit(prices []int) int {
	minprice := math.MaxInt32
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}

	}
	return maxprofit
}

func maxProfit1(prices []int) int {

	var sell float64 = 0
	var buy float64 = math.MaxFloat64
	for i := 0; i < len(prices); i++ {
		buy = math.Min(float64(prices[i]), buy)
		sell = math.Max(float64(prices[i])-buy, sell)
	}
	return int(sell)
}
