package main

import "fmt"

// map lesson

func main() {
	// stocks := map[string]float64{}
	// stocks["MSTF"] = 245.71
	// stocks["AAPL"] = 123.74
	// stocks["GOOGL"] = 2347.58

	// fmt.Println(stocks["APPL"])
	// fmt.Println(stocks{"FB"})

	stocks := map[string]float64{
		"MSFT":  245.71,
		"AAPL":  123.74,
		"GOOGL": 2347.58,
	}
	// fmt.Println(stocks)

	v, found := stocks["FB"]
	if found {
		fmt.Println("Stock value is", v)
	} else {
		fmt.Println("No data found")
	}

	// 値の削除方法
	delete(stocks, "MSFT")
	fmt.Println(stocks)
}
