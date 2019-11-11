package fxcm

import "encoding/json"

type CandlePeriod string

var (
	CandlePeriodMinute1  CandlePeriod = "m1"
	CandlePeriodMinute5  CandlePeriod = "m5"
	CandlePeriodMinute15 CandlePeriod = "m15"
	CandlePeriodMinute30 CandlePeriod = "m30"
	CandlePeriodHours1   CandlePeriod = "H1"
	CandlePeriodHours2   CandlePeriod = "H2"
	CandlePeriodHours3   CandlePeriod = "H3"
	CandlePeriodHours4   CandlePeriod = "H4"
	CandlePeriodHours6   CandlePeriod = "H6"
	CandlePeriodHours8   CandlePeriod = "H8"
	CandlePeriodDay      CandlePeriod = "D1"
	CandlePeriodWeek     CandlePeriod = "W1"
	CandlePeriodMonth    CandlePeriod = "M1"
)

type Candle struct {
	Timestamp uint32  `json:"timestamp"`
	BidOpen   float64 `json:"bid_open"`
	BidClose  float64 `json:"bid_close"`
	BidHigh   float64 `json:"bid_high"`
	BidLow    float64 `json:"bid_low"`
	AskOpen   float64 `json:"ask_open"`
	AskClose  float64 `json:"ask_close"`
	AskHigh   float64 `json:"ask_high"`
	AskLow    float64 `json:"ask_low"`
	TickQty   float64 `json:"tick_qty"`
}

func (t *Candle) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t Candle) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t Candle) GetGraphTmestamp() float64 {
	return float64(t.Timestamp) * 1000
}
