package response

type Unsubscribe struct {
	Header
	Pairs string `json:"pairs"`
}
