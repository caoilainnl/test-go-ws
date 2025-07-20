package main

import (
	"fmt"
	"time"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func TestHandleOrderBook() {
	bitget := ccxt.NewBitgetWs(map[string]interface{}{
		// "apiKey":    "your-api-key",
		// "secret":    "your-secret",
		// "password":  "your-password",
		// "sandbox":   true,
	})

	client := map[string]interface{}{
		"resolve": func(data interface{}, messageHash string) {
			fmt.Printf("Resolved: %s %v\n", messageHash, data)
		},
		"reject": func(error interface{}, messageHash string) {
			fmt.Printf("Rejected: %s %v\n", messageHash, error)
		},
		"subscriptions": map[string]interface{}{},
		"connected":     true,
		"disconnected":  false,
		"futures":       map[string]interface{}{},
		"rejections":    map[string]interface{}{},
	}

	message := map[string]interface{}{
		"action": "update",
		"arg": map[string]interface{}{
			"channel":  "books",
			"instId":   "XRPUSDT",
			"instType": "SPOT",
		},
		"data": []interface{}{
			map[string]interface{}{
				"asks": [][]interface{}{
					{3.4297, 921.3269},
					{3.4324, 8406.0000},
					{3.4326, 8213.0000},
					{3.4327, 11155.5000},
					{3.4337, 21700.5000},
					{3.4353, 16472.0000},
					{3.4402, 3150.8906},
					{3.4406, 3997.3984},
					{3.4410, 2962.7777},
					{3.4414, 3339.0034},
					{3.4418, 4185.5113},
					{3.4422, 3715.2291},
					{3.4426, 1928.1569},
					{3.4430, 2774.6648},
					{3.4438, 3527.1163},
					{3.4442, 1645.9876},
					{3.4446, 3903.3420},
					{3.4450, 2868.7213},
					{3.4454, 3056.8341},
					{3.4458, 3809.2856},
					{3.4462, 1551.9312},
					{3.4466, 2492.4955},
					{3.4470, 2304.3827},
					{3.4474, 4279.5677},
					{3.4478, 1834.1005},
					{3.4482, 4091.4549},
					{3.4486, 3433.0599},
					{3.4490, 2210.3262},
					{3.4494, 2680.6084},
					{3.4498, 2022.2134},
					{3.4502, 3244.9470},
				},
				"bids": [][]interface{}{
					{3.4296, 1751.4453},
					{3.4293, 0.0},
					{3.4269, 55696.5000},
					{3.4268, 28655.2500},
					{3.4252, 27354.0000},
					{3.4242, 21357.5000},
					{3.4239, 11312.0000},
					{3.4236, 24557.5000},
					{3.4230, 11646.0000},
					{3.4191, 2586.5520},
					{3.4187, 2304.3827},
					{3.4183, 3433.0599},
					{3.4179, 4185.5113},
					{3.4171, 2210.3262},
					{3.4167, 3244.9470},
					{3.4163, 2398.4391},
					{3.4159, 3056.8341},
					{3.4155, 3527.1163},
					{3.4151, 4091.4549},
					{3.4147, 3809.2856},
					{3.4143, 1834.1005},
					{3.4139, 2962.7777},
					{3.4135, 3903.3420},
					{3.4131, 2680.6084},
					{3.4127, 3997.3984},
					{3.4123, 1551.9312},
					{3.4119, 2116.2698},
					{3.4115, 3339.0034},
					{3.4111, 1928.1569},
					{3.4107, 1645.9876},
					{3.4103, 2492.4955},
					{3.4099, 1740.0441},
					{3.4095, 2868.7213},
					{3.4091, 2022.2134},
					{3.4069, 6.7249},
				},
				"checksum": -1618681252,
				"seq":      1330394126235410432,
				"ts":       1752916597899,
			},
		},
		"ts": 1752916597901,
	}


	bitget.HandleOrderBook(client, message)

}


func TestMethods() {
	// ------------------- CORE -------------------
	// ex := ccxt.NewBinanceWsCore()
	apiKey := ""
	secret := ""
	password := ""
	ex := ccxt.NewBitgetWs(map[string]interface{}{
		// "apiKey": apiKey,
		// "secret": secret,
		// "password": password,
	})
	ex.Verbose = true
	ex.Init(nil)
	ex.LoadMarkets()
	// ex.ApiKey = apiKey
	// ex.Secret = secret
	// ex.Password = password

	// ------------------- CORE -------------------
	symbol   			  		:= "XRP/USDT"
	symbols   			  		:= []string{symbol}
	params 				  		:= map[string]interface{}{}
	condition 			  		:= true
	var	tickerCh 				ccxt.Ticker		// <-chan interface{}
	var	tickersCh 				ccxt.Tickers	// <-chan interface{}
	var	tradesCh 				[]ccxt.Trade	// <-chan interface{}
	var	tradesForSymbolsCh 		[]ccxt.Trade	// <-chan interface{}
	var	bidsAsksCh 				ccxt.Tickers	// <-chan interface{}
	var	ohlcvCh 				[]ccxt.OHLCV	// <-chan interface{}
	var	orderBookCh 			ccxt.IOrderBook	// <-chan interface{}
	var	orderBookForSymbolsCh 	ccxt.IOrderBook	// <-chan interface{}
	var	balanceCh 				ccxt.Balances	// <-chan interface{}
	var	myTradesCh 				[]ccxt.Trade	// <-chan interface{}
	var	ordersCh 				[]ccxt.Order	// <-chan interface{}
	var	positionsCh 			[]ccxt.Position	// <-chan interface{}

	fmt.Println(symbols, params, tickerCh, tickersCh, tradesCh, tradesForSymbolsCh, bidsAsksCh, ohlcvCh, orderBookCh, orderBookForSymbolsCh, balanceCh, myTradesCh, ordersCh, positionsCh)

	time.AfterFunc(30*time.Second, func() {
		condition = false
	})

	for condition {
		// tickerCh, _			  			= ex.WatchTicker(symbol)
		// tickersCh, _ 					= ex.WatchTickers(ccxt.WithWatchTickersSymbols(symbols))
		// tradesCh, _ 			  		= ex.WatchTrades(symbol)
		// tradesForSymbolsCh, _  	  		= ex.WatchTradesForSymbols(symbols)
		// bidsAsksCh, _ 			  		= ex.WatchBidsAsks(ccxt.WithWatchBidsAsksSymbols(symbols))
		// ohlcvCh, _    			  		= ex.WatchOHLCV(symbol)					// empty list
		// orderBookCh, _ 		  			= ex.WatchOrderBook(symbol)
		orderBookForSymbolsCh, _ 		= ex.WatchOrderBookForSymbols(symbols)
		// balanceCh, _  					= ex.WatchBalance(symbol)
		// myTradesCh, _ 					= ex.WatchMyTrades(ccxt.WithWatchMyTradesSymbol(symbol))
		// ordersCh, _ 					= ex.WatchOrders(ccxt.WithWatchOrdersSymbol(symbol))
		// positionsCh, _ 					= ex.WatchPositions(ccxt.WithWatchPositionsSymbols([]string{"XRP/USDT:USDT"}))

		// fmt.Printf("tickerCh: %+v\n", tickerCh)
		// fmt.Printf("tickersCh: %+v\n", tickersCh)
		// fmt.Printf("tradesCh: %+v\n", tradesCh)
		// fmt.Printf("tradesForSymbolsCh: %+v\n", tradesForSymbolsCh)
		// fmt.Printf("bidsAsksCh: %+v\n", bidsAsksCh)
		// fmt.Printf("ohlcvCh: %+v\n", ohlcvCh)
		// fmt.Printf("orderBookCh: %+v\n", orderBookCh)
		fmt.Printf("orderBookForSymbolsCh: %+v\n", orderBookForSymbolsCh)
		// fmt.Printf("balanceCh: %+v\n", balanceCh)
		// fmt.Printf("myTradesCh: %+v\n", myTradesCh)
		// fmt.Printf("ordersCh: %+v\n", ordersCh)
		// fmt.Printf("positionsCh: %+v\n", positionsCh)
	}

	// fmt.Println(<-ex.UnWatchTicker(symbol, params))
	// fmt.Println(<-ex.UnWatchOHLCV(symbol, params))
	// fmt.Println(<-ex.UnWatchOrderBook(symbol, params))
	// fmt.Println(<-ex.UnWatchTrades(symbol, params))

}

func main() {
	TestMethods()
}