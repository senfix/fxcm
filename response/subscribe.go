package response

import "github.com/senfix/fxcm"

type Subscribe struct {
	Header
	Pairs []fxcm.Subscribe `json:"pairs"`
}
