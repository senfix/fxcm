package fxcm

type RestAPI interface {

	//Market data subscription
	//After subscribing, market price updates will be pushed to the client via the socket
	//200, 401, 403, 404
	Subscribe(pairs string) (err error, response Subscribe)

	//Market data unsubscription
	//unsubscribe market data
	//200, 401, 403, 404
	Unsubscribe(pairs string) (err error, response Unsubscribe)

	//This command will request historical price data.
	//200, 401, 403, 404
	GetCandles(period CandlePeriod, offerId int, candlesAmount int) (err error, candles []Candle)
	//TODO
	//TradingSubscribe
	//TradingUnsubscribe
	//GetOffer
	//GetOpenPosition
	//GetClosedPosition
	//GetOrder
	//GetSummary
	//GetAccount
	//Open_trade
	//Close_trade
	//Change_order
	//Delete_order
	//Create_entry_order
	//Simple_oco
	//Edit_oco
	//Change_trade_stop_limit
	//Change_order_stop_limit
	//Close_all_for_symbol

}
