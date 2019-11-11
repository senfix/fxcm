package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	cerrors "github.com/senfix/fxcm/errors"

	"github.com/segmentio/objconv/json"
	"github.com/senfix/fxcm/response"

	"github.com/senfix/fxcm"
)

type rest struct {
	client   http.Client
	token    string
	socketId string
}

func NewClient() fxcm.RestAPI {
	return &rest{
		client: http.Client{
			Timeout: time.Second * 10,
		},
		socketId: "EWJQ4XcHrkuYSB6LADQB",
		token:    "d6f003eb1cb2bdb2cc14eff01bfb31e1b15a5c48",
	}
}

func (t *rest) Subscribe(pairs string) (err error, response fxcm.Subscribe) {
	panic("implement me")
}

func (t *rest) Unsubscribe(pairs string) (err error, response fxcm.Unsubscribe) {
	panic("implement me")
}

func (t *rest) GetCandles(period fxcm.CandlePeriod, offerId int, candlesAmount int) (err error, candles []fxcm.Candle) {
	uri := fmt.Sprintf("%v/%v/%v?num=%v", CANDLES_URI, offerId, period, candlesAmount)
	request, _ := http.NewRequest("GET", uri, nil)
	r, err := t.client.Do(request)
	if err != nil {
		return cerrors.HttpCallError(err), candles
	}

	raw, _ := ioutil.ReadAll(r.Body)

	data := response.Candles{}

	if err = json.Unmarshal(raw, &data); err != nil {
		return cerrors.UnmarshalError(err), candles
	}

	if err = validateResponse(raw); err != nil {
		return
	}

	candles = data.GetModel()
	return
}
