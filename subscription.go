package fxcm

type Subscribe struct {
	Update uint      `json:"update"`
	Rates  []float64 `json:"rates"`
	Symbol string    `json:"symbol"`
}

type Unsubscribe struct {
	Pairs string `json:"pairs"`
}
