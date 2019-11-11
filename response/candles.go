package response

import "github.com/senfix/fxcm"

type Candles struct {
	InstrumentId int               `json:"instrument_id"`
	PeriodId     fxcm.CandlePeriod `json:"period_id"`
	Candles      [][]float64       `json:"candles"`
}

func (t *Candles) GetModel() []fxcm.Candle {
	data := make([]fxcm.Candle, 0)
	for _, d := range t.Candles {
		if len(d) != 10 {
			continue
		}
		data = append(data, fxcm.Candle{
			Timestamp: uint32(d[0]),
			BidOpen:   d[1],
			BidClose:  d[2],
			BidHigh:   d[3],
			BidLow:    d[4],
			AskOpen:   d[5],
			AskClose:  d[6],
			AskHigh:   d[7],
			AskLow:    d[8],
			TickQty:   d[9],
		})
	}
	return data
}
