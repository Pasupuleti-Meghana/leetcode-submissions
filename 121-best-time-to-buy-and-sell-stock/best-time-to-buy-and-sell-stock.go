func maxProfit(prices []int) int {
    minprice := prices[0]
    maxprofit := 0
    for _, price := range prices {
        if price < minprice {
            minprice = price
        } else if price-minprice > maxprofit {
            maxprofit = price - minprice
        }
    }
    return maxprofit
}


