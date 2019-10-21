package response

type Subscribe struct {
	Header
	Pairs SubscribePairs `json:"pairs"`
}

type SubscribePairs struct {
	Update uint      `json:"update"`
	Rates  []float64 `json:"rates"`
	Symbol string    `json:"symbol"`
}
