package main

import (
	"fmt"
	"time"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {
	// ------------------- CORE -------------------
	// ex := ccxt.NewBinanceWsCore()
	apiKey := "..."
	secret := "..."
	password := "..."
	ex := ccxt.NewBitgetWs(map[string]interface{}{
		"apiKey": apiKey,
		"secret": secret,
		"password": password,
	})
	ex.Init(nil)
	ex.LoadMarkets()
	ex.ApiKey = apiKey
	ex.Secret = secret
	ex.Password = password

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

	time.AfterFunc(7*time.Second, func() {
		condition = false
	})

	for condition {
		tickerCh, _			  			= ex.WatchTicker(symbol)
		// tickersCh, _ 					= ex.WatchTickers(ccxt.WithWatchTickersSymbols(symbols))
		// tradesCh, _ 			  		= ex.WatchTrades(symbol)
		// tradesForSymbolsCh, _ 	  		= ex.WatchTradesForSymbols(symbols)
		// bidsAsksCh, _ 			  		= ex.WatchBidsAsks(ccxt.WithWatchBidsAsksSymbols(symbols))
		// ohlcvCh, _    			  		= ex.WatchOHLCV(symbol)					// empty list
		// orderBookCh, _ 		  			= ex.WatchOrderBook(symbol)
		// orderBookForSymbolsCh, _ 		= ex.WatchOrderBookForSymbols(symbols)
		// balanceCh, _  					= ex.WatchBalance(symbol)
		// myTradesCh, _ 					= ex.WatchMyTrades(ccxt.WithWatchMyTradesSymbol(symbol))
		// ordersCh, _ 					= ex.WatchOrders(ccxt.WithWatchOrdersSymbol(symbol))
		// positionsCh, _ 					= ex.WatchPositions(ccxt.WithWatchPositionsSymbols([]string{"XRP/USDT:USDT"}))

		fmt.Printf("tickerCh: %+v\n", tickerCh)
		// fmt.Printf("tickersCh: %+v\n", tickersCh)
		// fmt.Printf("tradesCh: %+v\n", tradesCh)
		// fmt.Printf("tradesForSymbolsCh: %+v\n", tradesForSymbolsCh)
		// fmt.Printf("bidsAsksCh: %+v\n", bidsAsksCh)
		// fmt.Printf("ohlcvCh: %+v\n", ohlcvCh)
		// fmt.Printf("orderBookCh: %+v\n", orderBookCh)
		// fmt.Printf("orderBookForSymbolsCh: %+v\n", orderBookForSymbolsCh)
		// fmt.Printf("balanceCh: %+v\n", balanceCh)
		// fmt.Printf("myTradesCh: %+v\n", myTradesCh)
		// fmt.Printf("ordersCh: %+v\n", ordersCh)
		// fmt.Printf("positionsCh: %+v\n", positionsCh)
	}

	fmt.Println(<-ex.UnWatchTicker(symbol, params))
	// fmt.Println(<-ex.UnWatchOHLCV(symbol, params))
	// fmt.Println(<-ex.UnWatchOrderBook(symbol, params))
	// fmt.Println(<-ex.UnWatchTrades(symbol, params))

}